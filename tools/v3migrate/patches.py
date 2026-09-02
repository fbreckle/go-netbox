"""Exceptions and additions the conversion cannot derive from the v3 document.

Everything here is a divergence, so everything here carries a reason. The rule is that the v3
document decides what the spec contains, and the only admissible reason to differ is a defect in
that document.
"""

# The v2 root. `swagger`, `basePath`, `schemes` and `securityDefinitions` have no 1:1 OpenAPI 3
# translation and the client's auth wiring depends on this shape. `info` is not here: it is
# copied from the v3 document.
ROOT = {   'basePath': '/api',
    'consumes': ['application/json'],
    'host': 'localhost:8001',
    'produces': ['application/json'],
    'schemes': ['http'],
    'security': [{'Bearer': []}],
    'securityDefinitions': {   'Bearer': {   'in': 'header',
                                             'name': 'Authorization',
                                             'type': 'apiKey'}},
    'swagger': '2.0'}

# Paths deliberately not carried. Empty: every path the document serves is converted. Add an
# entry only with the reason it is left out.
EXCLUDED_PATHS: "dict[str, str]" = {}

# Properties the v3 document types less precisely than NetBox serves them. Each entry replaces
# the converted property wholesale, so it must say everything v3 says about it and nothing v3
# does not.
#
# `extra_choices` is a list of [value, label] pairs; v3 declares the pair as an array of two
# untyped items. Restoring `string` is the only thing added: the pair elements are strings
# (CustomFieldChoiceSet.extra_choices is an ArrayField of two-element string arrays, each
# element up to 100 characters).
_CHOICE_PAIR = {'type': 'array',
                'items': {'type': 'array', 'minItems': 2, 'maxItems': 2,
                          'items': {'type': 'string', 'maxLength': 100}}}
TYPE_PATCHES = {
    'CustomFieldChoiceSet.extra_choices': _CHOICE_PAIR,
    'WritableCustomFieldChoiceSet.extra_choices': _CHOICE_PAIR,
}
