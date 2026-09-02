"""Convert a NetBox OpenAPI 3 spec into this repo's OpenAPI 2 spec.

NetBox publishes OpenAPI 3 since 3.5, but go-swagger (which generates this client)
only reads OpenAPI 2.  The v3 document is NetBox's own account of its behaviour, so
it decides what this spec contains: the conversion walks its paths, emits every
operation, and pulls in the schemas they reference.  Nothing is inferred from an
object type, so a singleton, an action route and a collection are all just paths.

  python3 tools/v3migrate/migrate.py --v3 'NetBox REST API (4.6).yaml'   # -> swagger.processed.json

The only things not derived are in patches.py, each with its reason: the v2 root,
the excluded paths (none today), and the properties the v3 document types less
precisely than NetBox serves them.  A local correction is a workaround for a defect in that
document, never a preference.

Output is sorted, so the same input always converts to the same file.

Translation rules applied (v3 -> v2) are documented in tools/v3migrate/README.md.
"""
import argparse, collections, copy, json, os, sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from patches import ROOT, EXCLUDED_PATHS, TYPE_PATCHES


METHODS = ('get', 'post', 'put', 'patch', 'delete')
BULK_METHODS = {'put', 'patch', 'delete'}   # on a collection path: NetBox GUI multi-select actions
QUERY_SKIP = {'brief', 'fields', 'omit'}    # change the response shape; break typed decoding
SCALARS = ('integer', 'string', 'number', 'boolean')


def load_spec(path):
    with open(path) as fh:
        head = fh.read(64)
    if head.lstrip().startswith('{'):
        return json.load(open(path))
    import yaml
    try:
        from yaml import CSafeLoader as Loader
    except ImportError:
        from yaml import SafeLoader as Loader
    return yaml.load(open(path), Loader=Loader)


class Migrator:
    def __init__(self, v3):
        self.v3 = v3
        self.S3 = v3['components']['schemas']
        self.P3 = v3['paths']
        self.D = {}
        self.pending_nested = set()

    # ---------- v3 schema references ----------
    @staticmethod
    def _ref_name(s):
        return s['$ref'].split('/')[-1] if isinstance(s, dict) and '$ref' in s else None

    @staticmethod
    def map_ref(name):
        """v3 schema name -> v2 definition name: the identity, bar two drf-spectacular suffixes.

        This spec references whatever the v3 document references at each site. NetBox uses both
        Brief* and Nested* and they are not synonyms -- BriefTag and NestedTag are different
        schemas, and `tags` references the latter -- so translating one family into the other
        would silently pick a shape NetBox never serves.
        """
        if name.endswith('Request'):
            name = name[:-len('Request')]
        # drf-spectacular leaks the serializer class name into a few schemas
        # (BriefCircuitGroupAssignmentSerializer_); the object type is the part before it.
        if name.endswith('Serializer_'):
            name = name[:-len('Serializer_')]
        # OpenAPI 3 gives PATCH its own all-optional schema; OpenAPI 2 has one body definition per
        # operation shape, and this spec has always sent PUT and PATCH through the same model.
        if name.startswith('Patched'):
            name = name[len('Patched'):]
        return name

    def _fk_target(self, p):
        """Brief*/Nested* schema name if p is a foreign-key wrapper, else None."""
        if not isinstance(p, dict):
            return None
        r = self._ref_name(p)
        if r and (r.startswith('Brief') or r.startswith('Nested')):
            return r
        for key in ('allOf', 'oneOf', 'anyOf'):
            for sub in p.get(key) or []:
                got = self._fk_target(sub)
                if got:
                    return got
        return None

    def _want_definition(self, v3name):
        target = self.map_ref(v3name)
        if target not in self.D:
            self.pending_nested.add((target, v3name))
        return target

    def _inline_ref(self, v3ref):
        """The v3 schema a $ref points at when it is not object-shaped, else None.

        OpenAPI 3 gives a name to shapes OpenAPI 2 spells out inline: IntegerRange is an
        array of two integers, not an object. Turning one into a definition would emit a
        model with no fields, so the target is converted in place instead.
        """
        target = self.S3.get(v3ref.split('/')[-1])
        if not isinstance(target, dict):
            return None
        if target.get('type') == 'object' or 'properties' in target:
            return None
        return target

    @staticmethod
    def _pick_member(subs):
        """The member of an allOf/oneOf/anyOf this spec carries, or None.

        OpenAPI 2 has no combinators. NetBox uses them for two things: a lone `allOf: [$ref]` to
        attach `nullable` or a description to a reference, which is carried as the reference; and
        `oneOf` between a scalar and something else -- a foreign key given as an id or as an object
        on writes, and a string that is either valid or empty (`color`, `email`, `dns_name`:
        `oneOf: [{string, pattern}, {string, maxLength: 0}]`) -- where the first scalar member is
        the one carrying the constraint NetBox applies to a value.
        """
        subs = [x for x in subs if isinstance(x, dict)]
        scalar = [x for x in subs if x.get('type') in SCALARS and '$ref' not in x]
        if scalar:
            return scalar[0]
        return subs[0] if len(subs) == 1 else None

    # ---------- schema conversion ----------
    @staticmethod
    def clean_enum(out):
        """OpenAPI 3 puts null in an enum to express nullability; Swagger 2 uses x-nullable.

        go-swagger's validator template assumes every enum value matches the declared
        type, so a nil in a string enum aborts generation with
        "wrong type for value; expected string; got interface {}".
        NetBox also omits `type` when null sorts first in the enum, so infer it.
        """
        values = out.get('enum')
        if not isinstance(values, list):
            return out
        if None in values:
            values = [v for v in values if v is not None]
            out['enum'] = values
            out['x-nullable'] = True
        if values and 'type' not in out and '$ref' not in out:
            kinds = set(map(type, values))
            if kinds == {str}:
                out['type'] = 'string'
            elif kinds == {int}:
                out['type'] = 'integer'
            elif kinds == {bool}:
                out['type'] = 'boolean'
        return out

    # go-swagger renders numeric bounds as untyped float constants, so an int64 bound past
    # 2**53 does not survive the round trip: NetBox's "maximum: 9223372036854775807" comes out
    # as 9.223372036854776e+18, which the generated validate.MaximumInt call cannot pass as an
    # int64 and the models package stops compiling. Such a bound only restates the Go type, so
    # drop it.
    INT_BOUND_LIMIT = 2 ** 53

    @classmethod
    def clean_bounds(cls, out):
        if out.get('type') != 'integer':
            return out
        for key in ('maximum', 'minimum'):
            value = out.get(key)
            if isinstance(value, int) and not isinstance(value, bool) and abs(value) > cls.INT_BOUND_LIMIT:
                del out[key]
        return out

    def conv_read(self, p, depth=0):
        if not isinstance(p, dict) or depth > 8:
            return copy.deepcopy(p)
        fk = self._fk_target(p)
        if fk:
            return {'$ref': '#/definitions/' + self._want_definition(fk)}
        out = {}
        for k, v in p.items():
            if k == 'x-spec-enum-id':
                continue
            if k == 'nullable':
                if v:
                    out['x-nullable'] = True
                continue
            if k == '$ref':
                inline = self._inline_ref(v)
                if inline is not None:
                    for mk, mv in self.conv_read(inline, depth + 1).items():
                        out.setdefault(mk, mv)
                    continue
                out['$ref'] = '#/definitions/' + self._want_definition(v.split('/')[-1])
                continue
            if k == 'properties':
                out['properties'] = {pk: self.conv_read(pv, depth + 1) for pk, pv in v.items()}
                continue
            if k == 'items':
                out['items'] = self.conv_read(v, depth + 1)
                continue
            if k in ('allOf', 'oneOf', 'anyOf'):
                pick = self._pick_member(v)
                if pick is not None:
                    for mk, mv in self.conv_read(pick, depth + 1).items():
                        out.setdefault(mk, mv)
                continue
            out[k] = copy.deepcopy(v)
        return self.clean_bounds(self.clean_enum(out))

    # NetBox references NestedTag (not BriefTag) at every `tags` site, so this agrees with the v3
    # document rather than overriding it. It stays load-bearing: without the short-circuit below,
    # _fk_target would see the Nested prefix on NestedTagRequest and turn `tags` into a list of ids.
    TAGS = {'type': 'array', 'items': {'$ref': '#/definitions/NestedTag'}}

    def conv_write(self, key, p, depth=0):
        # ConfigContext.tags is a list of slugs, not of tag objects, so the override is conditional
        # on v3 actually referencing a tag schema here.
        if key == 'tags' and self._fk_target((p or {}).get('items') or {}):
            return copy.deepcopy(self.TAGS)
        if not isinstance(p, dict):
            return copy.deepcopy(p)
        nullable = bool(p.get('nullable'))
        if self._fk_target(p):
            out = {'type': 'integer'}
            if nullable:
                out['x-nullable'] = True
            return out
        if p.get('type') == 'array' and self._fk_target(p.get('items') or {}):
            out = {'type': 'array', 'items': {'type': 'integer'}, 'uniqueItems': True}
            if nullable:
                out['x-nullable'] = True
            return out
        out = {}
        for k, v in p.items():
            if k == 'x-spec-enum-id':
                continue
            if k == 'nullable':
                if v:
                    out['x-nullable'] = True
                continue
            if k == '$ref':
                inline = self._inline_ref(v)
                if inline is not None:
                    for mk, mv in self.conv_write(key, inline, depth + 1).items():
                        out.setdefault(mk, mv)
                    continue
                out['$ref'] = '#/definitions/' + self._want_definition(v.split('/')[-1])
                continue
            if k == 'items':
                out['items'] = self.conv_write(None, v, depth + 1)
                continue
            if k == 'properties':
                out['properties'] = {pk: self.conv_write(pk, pv, depth + 1) for pk, pv in v.items()}
                continue
            if k in ('allOf', 'oneOf', 'anyOf'):
                pick = self._pick_member(v)
                if pick is not None:
                    for mk, mv in self.conv_write(key, pick, depth + 1).items():
                        out.setdefault(mk, mv)
                continue
            out[k] = copy.deepcopy(v)
        return self.clean_bounds(self.clean_enum(out))

    # ---------- definitions ----------
    def build_read(self, name, source=None):
        s = self.S3[source or name]
        props = s.get('properties') or {}
        read_only = {k for k, v in props.items() if v.get('readOnly')}
        out = {'type': 'object'}
        req = sorted(r for r in (s.get('required') or []) if r not in read_only)
        if req:
            out['required'] = req
        out['properties'] = {k: self.conv_read(v) for k, v in props.items()}
        return out

    def write_source(self, name):
        for cand in ('Writable%sRequest' % name, '%sRequest' % name):
            if cand in self.S3:
                return cand
        return None

    def needs_writable(self, name):
        """Whether writes to `name` need their own definition.

        True whenever the converted write shape of any property differs from its converted read
        shape, or the write schema has a field the read schema does not. OpenAPI 3 keeps the two
        apart; OpenAPI 2 needs a second definition to say the same thing. A narrower test that only
        looked for foreign keys missed NotificationGroup, whose `groups` reads as nested objects and
        writes as ids, and User, which accepts a password it never returns.
        """
        src = self.write_source(name)
        if not src:
            return False
        if src.startswith('Writable'):
            return True
        request = self.S3[src].get('properties') or {}
        read = (self.S3.get(name) or {}).get('properties') or {}
        if read and set(request) - set(read):
            return True
        for key, value in request.items():
            if key not in read:
                continue
            here = json.dumps(self.conv_write(key, value), sort_keys=True)
            there = json.dumps(self.conv_read(read[key]), sort_keys=True)
            if here != there:
                return True
        return False

    def build_write(self, name, source=None):
        s = self.S3[source or self.write_source(name)]
        out = {'type': 'object'}
        if s.get('required'):
            out['required'] = sorted(s['required'])
        out['properties'] = {k: self.conv_write(k, v) for k, v in (s.get('properties') or {}).items()}
        return out

    DEFAULT_RESPONSE = {'description': '', 'schema': {'type': 'object', 'additionalProperties': True}}

    # ---------- operations ----------
    def query_params(self, op):
        return [self._query_param(prm) for prm in op.get('parameters') or []
                if prm.get('in') == 'query' and prm['name'] not in QUERY_SKIP]

    @staticmethod
    def _query_param(prm):
        """A v2 query parameter carrying the type the v3 document declares.

        Swagger 2 puts type/items/collectionFormat inline on the parameter -- there is no schema
        object for a non-body parameter. NetBox declares its repeatable filters `style: form,
        explode: true`, which is `collectionFormat: multi`: `?id=1&id=2`, not `?id=1,2`. Its
        filters are OR, so repeating one widens the match.

        The integer `format` is dropped so ids come out as int64 like every other id in the client,
        and an `enum` is dropped because go-swagger would validate it client-side and reject a
        value NetBox accepts.
        """
        schema = prm.get('schema') or {}

        def scalar(sch):
            return {'type': sch.get('type') or 'string'}

        out = {'in': 'query', 'name': prm['name'], 'required': False}
        if prm.get('description'):
            out['description'] = prm['description']
        if schema.get('type') == 'array':
            out.update({'type': 'array', 'items': scalar(schema.get('items') or {}),
                        'collectionFormat': 'multi'})
        else:
            out.update(scalar(schema))
        return out

    def paginated(self, name):
        return {'required': ['count', 'results'], 'type': 'object', 'properties': {
            'count': {'type': 'integer'},
            'next': {'type': 'string', 'format': 'uri', 'x-nullable': True},
            'previous': {'type': 'string', 'format': 'uri', 'x-nullable': True},
            'results': {'type': 'array', 'items': {'$ref': '#/definitions/' + name}}}}

    @staticmethod
    def _body_schema(op):
        try:
            return op['requestBody']['content']['application/json']['schema']
        except (KeyError, TypeError):
            return None

    @staticmethod
    def _response_schema(op):
        for code in ('200', '201'):
            try:
                return code, op['responses'][code]['content']['application/json']['schema']
            except (KeyError, TypeError):
                continue
        return None, None

    @staticmethod
    def _is_keyed(q):
        return any(seg.startswith('{') for seg in q.strip('/').split('/'))

    def _write_ref(self, v3name):
        """The v2 $ref for a v3 request schema: a Writable twin when the write shape differs."""
        base = self.map_ref(v3name)
        if not self.needs_writable(base):
            return {'$ref': '#/definitions/' + self._want_definition(v3name)}
        # v3 already spells some of them Writable<X>Request; do not double the prefix
        target = base if base.startswith('Writable') else 'Writable' + base
        if target not in self.D:
            self.D[target] = self.build_write(base)
        return {'$ref': '#/definitions/' + target}

    def _body_payload(self, body):
        """Convert a request-body schema.

        NetBox's create endpoints accept one object or a list of them, which v3 spells
        `oneOf: [$ref X, array of $ref X]`. OpenAPI 2 has one body per operation and this client
        sends one object, so the object member is carried; the list form is the bulk create the
        GUI uses, left out with the other bulk operations.
        """
        if 'oneOf' in body:
            refs = [s for s in body['oneOf'] if isinstance(s, dict) and '$ref' in s]
            arrays = [s for s in body['oneOf'] if isinstance(s, dict) and s.get('type') == 'array']
            if len(refs) == 1 and len(arrays) == 1 \
                    and (arrays[0].get('items') or {}).get('$ref') == refs[0]['$ref']:
                body = refs[0]
        name = body.get('$ref', '').split('/')[-1]
        if name:
            return self._write_ref(name)
        if body.get('type') == 'array' and (body.get('items') or {}).get('$ref'):
            return {'type': 'array', 'items': self._write_ref(body['items']['$ref'].split('/')[-1])}
        return self.conv_write(None, body)

    def _response_payload(self, schema):
        """Convert a response schema, collapsing NetBox's Paginated*List into this spec's shape."""
        if not isinstance(schema, dict):
            return None
        name = schema.get('$ref', '').split('/')[-1]
        if name.startswith('Paginated') and name.endswith('List'):
            items = ((self.S3[name].get('properties') or {}).get('results') or {}).get('items') or {}
            inner = items.get('$ref', '').split('/')[-1]
            if inner:
                return self.paginated(self._want_definition(inner))
        if name:
            return {'$ref': '#/definitions/' + self._want_definition(name)}
        if schema.get('type') == 'array':
            item = self._response_payload(schema.get('items') or {})
            return {'type': 'array', 'items': item} if item else None
        return self.conv_read(schema)

    def convert_operation(self, q, meth, op):
        params = []
        for prm in op.get('parameters') or []:
            if prm.get('in') != 'path':
                continue
            params.append({'name': prm['name'], 'in': 'path', 'required': True,
                           'description': prm.get('description', '') or '',
                           'type': (prm.get('schema') or {}).get('type', 'string')})
        params += self.query_params(op)
        body = self._body_schema(op)
        if body is not None:
            params.append({'name': 'data', 'in': 'body', 'required': True,
                           'schema': self._body_payload(body)})

        responses = {'default': copy.deepcopy(self.DEFAULT_RESPONSE)}
        code, resp = self._response_schema(op)
        if code is None:
            code = '204' if meth == 'delete' else '200'
            responses[code] = {'description': ''}
        else:
            conv = self._response_payload(resp)
            responses[code] = {'description': ''} if conv is None else {'description': '', 'schema': conv}
        # NetBox's operation ids are unique across the whole document, so they are used verbatim:
        # deriving one would be this spec inventing a name NetBox already gave.
        return {'operationId': op['operationId'], 'description': op.get('description', '') or '',
                'tags': [q.strip('/').split('/')[0]], 'parameters': params, 'responses': responses}

    # ---------- the conversion ----------
    def convert(self):
        """Build a complete v2 spec from the v3 document, path by path.

        Everything this spec contains is reachable from a v3 path: the operations it declares and,
        transitively, the schemas they reference. Nothing is inferred from an object-type shape, so
        a singleton, an action route and an ordinary collection are all just paths.
        """
        v2 = copy.deepcopy(ROOT)
        v2['info'] = copy.deepcopy(self.v3['info'])
        v2['paths'], v2['definitions'] = {}, {}
        self.D, self.pending_nested = v2['definitions'], set()

        for v3path in sorted(self.P3):
            q = v3path[4:] if v3path.startswith('/api') else v3path
            if q in EXCLUDED_PATHS:
                continue
            keyed = self._is_keyed(q)
            entry = {}
            for meth in METHODS:
                op = self.P3[v3path].get(meth)
                if not isinstance(op, dict):
                    continue
                if meth in BULK_METHODS and not keyed:
                    continue  # NetBox's GUI multi-select actions; this client has never had them
                entry[meth] = self.convert_operation(q, meth, op)
            if entry:
                v2['paths'][q] = entry
        stuck = self.resolve_nested()
        # Sort both maps so the same input always produces the same file, byte for byte: the
        # definitions are discovered by walking paths, and discovery order is not meaningful.
        v2['paths'] = dict(sorted(v2['paths'].items()))
        v2['definitions'] = dict(sorted(v2['definitions'].items()))
        self.D = v2['definitions']
        for name, patch in TYPE_PATCHES.items():
            defn = self.D.get(name.split('.')[0])
            if defn and name.split('.')[1] in (defn.get('properties') or {}):
                defn['properties'][name.split('.')[1]] = copy.deepcopy(patch)
        return v2, stuck

    def resolve_nested(self, rounds=64):
        """Build every definition a converted schema referenced, until nothing new is referenced.

        A reference is usually to a read schema. NetBox also declares a few request-only shapes --
        PrefixLengthRequest, ScriptInputRequest -- with no read twin; those are built from the
        request schema. A name found in neither place is a defect in the v3 document, reported by
        the caller rather than papered over.
        """
        for _ in range(rounds):
            todo = sorted((n, src) for n, src in self.pending_nested if n not in self.D)
            if not todo:
                return []
            for target, v3name in todo:
                base = v3name[:-7] if v3name.endswith('Request') else v3name
                if base in self.S3:
                    self.D[target] = self.build_read(target, source=base)
                elif v3name in self.S3:
                    self.D[target] = self.build_write(target, source=v3name)
        return [n for n, _ in self.pending_nested if n not in self.D]


PY_TYPES = {'string': str, 'integer': int, 'number': (int, float), 'boolean': bool}


def bad_enums(name, node, path=''):
    """Enum values that do not match the declared type abort go-swagger's templates."""
    out = []
    if isinstance(node, dict):
        declared = node.get('type')
        declared = declared if isinstance(declared, str) else None
        values = node.get('enum')
        if isinstance(values, list):
            if None in values:
                out.append('%s%s: enum contains null (use x-nullable instead)' % (name, path))
            if declared is None:
                out.append('%s%s: enum without a type' % (name, path))
            elif declared in PY_TYPES:
                wrong = [v for v in values if not isinstance(v, PY_TYPES[declared])]
                if wrong:
                    out.append('%s%s: enum values %r are not %s' % (name, path, wrong[:3], declared))
        for key, value in node.items():
            if key in ('properties',):
                for pk, pv in (value or {}).items():
                    out.extend(bad_enums(name, pv, path + '/' + pk))
            elif key in ('items', 'additionalProperties') and isinstance(value, dict):
                out.extend(bad_enums(name, value, path + '/' + key))
            elif key == 'allOf':
                for i, sub in enumerate(value):
                    out.extend(bad_enums(name, sub, path + '/allOf%d' % i))
    return out


def bad_bounds(name, node, path=''):
    """Integer bounds go-swagger cannot render: it emits them as float64 constants, so a bound
    past 2**53 does not fit the int64 field it guards and the package stops compiling."""
    out = []
    if isinstance(node, dict):
        if node.get('type') == 'integer':
            for key in ('maximum', 'minimum'):
                value = node.get(key)
                if isinstance(value, int) and not isinstance(value, bool) and abs(value) > Migrator.INT_BOUND_LIMIT:
                    out.append('%s%s: %s %d does not survive go-swagger\'s float64 rendering'
                               % (name, path, key, value))
        for key, value in node.items():
            if key in ('properties',):
                for pk, pv in (value or {}).items():
                    out.extend(bad_bounds(name, pv, path + '/' + pk))
            elif key in ('items', 'additionalProperties') and isinstance(value, dict):
                out.extend(bad_bounds(name, value, path + '/' + key))
            elif key == 'allOf':
                for i, sub in enumerate(value):
                    out.extend(bad_bounds(name, sub, path + '/allOf%d' % i))
    return out


def readonly_required(definition):
    """required entries that are readOnly -- go-swagger warns on these."""
    props = definition.get('properties') or {}
    return sorted(r for r in (definition.get('required') or []) if (props.get(r) or {}).get('readOnly'))


def structural_problems(spec):
    """Faults that would break go-swagger or the generated client, whatever produced the spec."""
    problems = []
    refs = {r for r in json.dumps(spec).split('"') if r.startswith('#/definitions/')}
    for ref in sorted(refs):
        if ref.split('/')[-1] not in spec['definitions']:
            problems.append('unresolved $ref: ' + ref)
    seen = collections.Counter(op['operationId'] for ops in spec['paths'].values()
                               for meth, op in ops.items()
                               if meth in METHODS and isinstance(op, dict))
    problems += ['duplicate operationId: %s (x%d)' % (k, n) for k, n in sorted(seen.items()) if n > 1]
    for name, defn in sorted(spec['definitions'].items()):
        problems += bad_enums(name, defn) + bad_bounds(name, defn) + readonly_required(defn)
    return problems


def summarise(spec):
    ops = sum(1 for ops in spec['paths'].values() for m in ops if m in METHODS)
    return '%d paths, %d operations, %d definitions' % (len(spec['paths']), ops,
                                                        len(spec['definitions']))


def main():
    ap = argparse.ArgumentParser(description=__doc__,
                                 formatter_class=argparse.RawDescriptionHelpFormatter)
    ap.add_argument('--v3', required=True, help='NetBox OpenAPI 3 spec (.yaml or .json)')
    ap.add_argument('--out', default='swagger.processed.json', help='the OpenAPI 2 spec to write')
    args = ap.parse_args()

    built, stuck = Migrator(load_spec(args.v3)).convert()
    problems = ['could not build: ' + n for n in sorted(stuck)] + structural_problems(built)
    if problems:
        print('the conversion is not sound (%d):' % len(problems), file=sys.stderr)
        for p in problems[:40]:
            print('   ' + p, file=sys.stderr)
        return 1

    with open(args.out, 'w') as fh:
        json.dump(built, fh, indent=2)
        fh.write('\n')
    print('converted %s -> %s\n  %s\n  %d exceptions applied from patches.py (%d typed properties, %d excluded paths)'
          % (args.v3, args.out, summarise(built),
             len(TYPE_PATCHES) + len(EXCLUDED_PATHS), len(TYPE_PATCHES), len(EXCLUDED_PATHS)))
    return 0


if __name__ == '__main__':
    sys.exit(main())
