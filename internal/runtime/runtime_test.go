package runtime

import (
	"context"
	"testing"

	"glass-box/internal/classifier"
	"glass-box/internal/constraints"
	"glass-box/internal/ledger"
	"glass-box/internal/state"
)

func TestExecutorProducesReceiptAndSignedLedgerRecord(t *testing.T) {
	initial := state.NewState("session-1", "actor-1", "request-1")
	executor := Executor{
		Model: StaticModel{Response: "answer", Reasoning: "visible reason"},
		Classifiers: []classifier.Classifier{
			classifier.StaticClassifier{
				ClassifierID: "static",
				Decisions: []classifier.Decision{{
					Name:       "risk",
					Category:   "general",
					Confidence: 0.1,
					Reason:     "low risk",
				}},
			},
		},
		Constraints: []constraints.Constraint{
			{Name: "requires_reasoning", Raw: `require reasoning != ""`, Severity: constraints.SeverityBlock},
		},
		Secret: []byte("test-secret"),
	}

	next, receipt, record, err := executor.Execute(context.Background(), initial, "prompt")
	if err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	if next.Model.Response != "answer" {
		t.Fatal("next state did not capture model response")
	}
	if len(receipt.ClassifierDecisions) != 1 || len(receipt.ConstraintResults) != 1 {
		t.Fatal("receipt did not capture classifier and constraint evidence")
	}
	if !ledger.VerifyRecord(record, []byte("test-secret")) {
		t.Fatal("expected transition ledger record to verify")
	}
}
