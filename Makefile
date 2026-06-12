.PHONY: test run

test:
	go test ./...

run:
	go run ./cmd/gbx-cli "hello glassbox"
