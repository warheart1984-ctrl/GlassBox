package classifier

import "testing"

func TestStaticClassifierReturnsTransparentDecision(t *testing.T) {
	cls := StaticClassifier{
		ClassifierID: "test-classifier",
		Decisions: []Decision{{
			Name:       "risk",
			Category:   "general",
			Confidence: 0.25,
			Reason:     "static test decision",
		}},
	}

	decisions, err := cls.Evaluate("hello", map[string]any{"actor": "tester"})
	if err != nil {
		t.Fatalf("Evaluate returned error: %v", err)
	}

	if cls.ID() != "test-classifier" || len(decisions) != 1 {
		t.Fatal("classifier did not expose expected identity and decisions")
	}
	if decisions[0].Reason == "" {
		t.Fatal("expected transparent decision reason")
	}
}
