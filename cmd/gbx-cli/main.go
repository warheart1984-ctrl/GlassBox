package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"glass-box/internal/constraints"
	"glass-box/internal/runtime"
	"glass-box/internal/state"
)

func main() {
	input := "hello glassbox"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	executor := runtime.Executor{
		Model: runtime.StaticModel{Response: "static governed response", Reasoning: "static model path"},
		Constraints: []constraints.Constraint{
			{Name: "requires_reasoning", Raw: `require reasoning != ""`, Severity: constraints.SeverityBlock},
		},
		Secret:        []byte("dev-secret"),
		PolicyVersion: "dev.v1",
	}

	_, receipt, record, err := executor.Execute(context.Background(), state.NewState("local", "cli", "request"), input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out := map[string]any{"receipt": receipt, "record": record}
	_ = json.NewEncoder(os.Stdout).Encode(out)
}
