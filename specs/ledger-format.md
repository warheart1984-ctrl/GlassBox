# Ledger Format

Ledger records are JSON-compatible objects with:

- `id`
- `type`
- `timestamp`
- `payload`
- `signature`

The MVP signs the canonical unsigned record fields using HMAC-SHA256.
Verification recomputes the signature and rejects tampered payloads.
