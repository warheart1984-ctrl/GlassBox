package constraints

import "testing"

func TestEvaluateBlockConstraintFailsWhenExpressionMatches(t *testing.T) {
	c := Constraint{Name: "no_external_http", Raw: `deny action == "http_request"`, Severity: SeverityBlock}

	result := Evaluate(c, map[string]any{"action": "http_request"})

	if result.Status != ResultFail {
		t.Fatalf("expected fail, got %s", result.Status)
	}
}

func TestEvaluateRequireConstraintPassesWhenExpressionMatches(t *testing.T) {
	c := Constraint{Name: "requires_reasoning", Raw: `require reasoning != ""`, Severity: SeverityBlock}

	result := Evaluate(c, map[string]any{"reasoning": "because"})

	if result.Status != ResultPass {
		t.Fatalf("expected pass, got %s", result.Status)
	}
}
