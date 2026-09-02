# v3migrate

Converts the NetBox **OpenAPI 3** document checked in at the repo root into the **OpenAPI 2**
`swagger.processed.json` the client is generated from. The v3 document is tracked; the v2 file is a
build intermediate, rebuilt by `make generate` and not tracked.

NetBox has published OpenAPI 3 since 3.5, but go-swagger — which generates this client — only
reads OpenAPI 2. The v3 document is NetBox's own account of its behaviour, so it decides what this
spec contains; the conversion is mechanical, and the only things that are not are in `patches.py`,
each with the reason it is there.

## Usage

```sh
make spec                # python3 tools/v3migrate/migrate.py --v3 'NetBox REST API (4.6).yaml'
make generate            # the above, then go-swagger
make check               # regenerate and fail if the tracked client changed
```

`make` picks the first `NetBox*.yaml` in the repo root; pass `v3spec=...` to point it elsewhere.
Moving to a newer NetBox means replacing that file with the document downloaded from
`/api/schema/` (it is YAML, whatever the file is named) and committing it with the regenerated
client.

Requires PyYAML (`python3 -m pip install --user pyyaml`) unless the spec is already JSON.

## How it works

It walks the v3 document's `paths` and emits every operation with **every** query parameter it
declares — lookup expansions included, so the client can express any filter NetBox accepts —
plus its body and responses, then pulls in the schemas those reference,
transitively. Nothing is inferred from an object type, so a singleton like `/status/`, an action
route like `/ipam/prefixes/{id}/available-ips/` and an ordinary collection are all just paths —
"which endpoints did we miss?" is not a question the conversion can get wrong.

Operation ids are NetBox's own, unique across the whole document (a duplicate aborts the
conversion), so there is nothing to derive or disambiguate, and the generated Go method names
follow from them. The `info` block is
NetBox's too.

The output is sorted and therefore byte-reproducible: the same v3 document always converts to the
same file, so a diff only ever shows a real change.

`make check` regenerates the client and fails if the tracked one differs. It is the whole drift
test: there are no separate categories to keep in step, because the conversion is the definition
of what the client should be.

The conversion refuses to write a spec with a `$ref` to a schema the v3 document does not declare, a
duplicate operation id, an enum containing `null` or holding values that do not match its type, an
integer bound beyond 2^53, or a `readOnly` property in a `required` list — faults that would break
go-swagger or the generated client whatever produced them.

## What is not derived

`patches.py`, and it is short:

- **`ROOT`** — the v2 root. `swagger`, `basePath`, `schemes` and `securityDefinitions` have no
  1:1 OpenAPI 3 translation and the client's auth wiring depends on this shape.
- **`TYPE_PATCHES`** — properties the v3 document types less precisely than NetBox serves them:
  `extra_choices` is a list of `[value, label]` string pairs, declared as pairs of untyped items.
  A patch replaces the converted property wholesale, so it must say everything v3 says and
  nothing v3 does not.
- **`EXCLUDED_PATHS`** — empty: every path the document serves is converted. Add an entry only
  with the reason it is left out.

## Translation rules

| OpenAPI 3 | OpenAPI 2 |
| --- | --- |
| `nullable: true` | `x-nullable: true` |
| `allOf: [$ref X]` + `nullable` (read) | `$ref: X` |
| `oneOf: [integer, Brief<X>Request]` (write) | `type: integer` |
| array of `Brief<X>Request` (write) | array of `integer`, `uniqueItems` |
| `oneOf: [{string, pattern}, {string, maxLength: 0}]` (valid or empty) | the constrained member |
| `tags` referencing a tag schema | array of `$ref NestedTag` |
| `<X>Request` / `Patched<X>Request` | `Writable<X>` when the write shape differs, else `<X>` |
| `<X>Request` with no `<X>` (`PrefixLengthRequest`, `ScriptInputRequest`) | `<X>`, built from the request schema |
| `Brief<X>Serializer_` | `Brief<X>` (drf-spectacular leaks the serializer class name) |
| create body `oneOf: [$ref X, array of $ref X]` (one or many) | `$ref X`; the list form is the GUI's bulk create |
| `readOnly` fields listed in `required` | dropped from `required` |
| `null` inside an `enum` | dropped, and `x-nullable: true` set |
| `enum` with no `type` | `type` inferred from the values |
| `x-spec-enum-id` | dropped |
| integer `maximum`/`minimum` beyond 2^53 | dropped (go-swagger renders them as float64) |
| `$ref` to a non-object schema (`IntegerRange`) | inlined, since OpenAPI 2 spells the shape out |
| `Paginated<X>List` | this spec's inline `{count, next, previous, results}` |
| bulk `PUT`/`PATCH`/`DELETE` on a collection | not emitted; they exist for NetBox's GUI |
| `brief` / `fields` / `omit` query parameters | skipped; they change the response shape and break typed decoding |
| query parameter `format` and `enum` | dropped: ids stay int64, and go-swagger would reject values NetBox accepts |

A `$ref` points at the schema NetBox points at, under the same name. NetBox uses both `Brief*`
and `Nested*` schemas and they are **not** synonyms — `BriefTag` and `NestedTag` differ,
and `tags` references the latter — so translating one family into the other would pick shapes
NetBox never serves.

A write definition carries exactly the properties its v3 request schema declares: no `id`, `url`
or `display`, which NetBox does not accept on a write. A property v3 leaves untyped (`{}`, the
polymorphic `assigned_object`, `link_peers`, `scope` and the JSON-blob fields) stays untyped and
becomes `interface{}`; that is what NetBox says about it.
