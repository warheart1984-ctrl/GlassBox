# Constraint DSL

The MVP supports two line forms:

```text
deny field == "value"
require field != ""
```

`deny` fails when the expression matches.
`require` fails when the expression does not match.

Supported operators in this build:

- `==`
- `!=`
