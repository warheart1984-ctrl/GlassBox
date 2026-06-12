# Glassbox

Glassbox is a transparent, auditable AI governance runtime prototype.

The current build is an MVP with:

- typed state transitions
- HMAC-signed ledger records
- transparent classifier decisions
- simple constraint evaluation
- runtime receipts
- a small CLI and operator ledger reader

## Run

```bash
go test ./...
go run ./cmd/gbx-cli "explain the system"
```

## Layout

```text
cmd/gbx-cli          demo execution CLI
cmd/gbx-operator     operator tools
internal/state       deterministic state and transitions
internal/ledger      signed audit records
internal/constraints constraint definitions and evaluator
internal/classifier  classifier interface and static implementation
internal/runtime     execution pipeline and receipts
internal/governance  policy structures and validation
specs/               governance design notes
```
