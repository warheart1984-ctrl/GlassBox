package state

import "testing"

func TestApplyTransitionBuildsNextState(t *testing.T) {
	initial := NewState("session-1", "actor-1", "request-1")

	next, transition := ApplyTransition(initial, "prompt", "input", "output")

	if next.ID == initial.ID {
		t.Fatal("expected a new state id")
	}
	if next.Version != initial.Version+1 {
		t.Fatalf("expected version %d, got %d", initial.Version+1, next.Version)
	}
	if transition.FromStateID != initial.ID || transition.ToStateID != next.ID {
		t.Fatal("transition did not preserve state lineage")
	}
	if next.Model.Prompt != "input" || next.Model.Response != "output" {
		t.Fatal("next state did not capture model input/output")
	}
}
