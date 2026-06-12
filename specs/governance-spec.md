# Governance Spec

Every governed execution must produce visible, constrained, reproducible, and auditable evidence.

The MVP pipeline is:

1. Start with a typed state.
2. Run transparent classifiers.
3. Execute a model adapter.
4. Evaluate constraints over input, response, reasoning, and action context.
5. Produce a receipt.
6. Sign a ledger transition record.

Blocked constraints stop the transition and return an error.
