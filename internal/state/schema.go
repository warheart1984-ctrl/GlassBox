package state

import "glass-box/internal/util"

type StateID string
type TransitionID string

type State struct {
	ID          StateID            `json:"id"`
	Version     int                `json:"version"`
	Context     map[string]any     `json:"context"`
	Model       ModelState         `json:"model"`
	Constraints ConstraintSnapshot `json:"constraints"`
	Provenance  Provenance         `json:"provenance"`
}

type ModelState struct {
	Prompt    string         `json:"prompt"`
	Response  string         `json:"response"`
	Reasoning string         `json:"reasoning"`
	Metadata  map[string]any `json:"metadata"`
}

type ConstraintSnapshot struct {
	Active  []string           `json:"active"`
	Results []ConstraintResult `json:"results"`
}

type ConstraintResult struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Details string `json:"details"`
}

type Provenance struct {
	SessionID    string `json:"session_id"`
	ActorID      string `json:"actor_id"`
	RequestID    string `json:"request_id"`
	ClassifierID string `json:"classifier_id"`
}

type Transition struct {
	ID          TransitionID `json:"id"`
	FromStateID StateID      `json:"from_state_id"`
	ToStateID   StateID      `json:"to_state_id"`
	Action      string       `json:"action"`
	Input       string       `json:"input"`
	Output      string       `json:"output"`
}

func NewState(sessionID, actorID, requestID string) State {
	return State{
		ID:      StateID(util.NewID("state")),
		Version: 1,
		Context: map[string]any{
			"session_id": sessionID,
			"actor_id":   actorID,
			"request_id": requestID,
		},
		Model: ModelState{
			Metadata: map[string]any{},
		},
		Constraints: ConstraintSnapshot{},
		Provenance: Provenance{
			SessionID: sessionID,
			ActorID:   actorID,
			RequestID: requestID,
		},
	}
}

func ApplyTransition(current State, action, input, output string) (State, Transition) {
	next := current
	next.ID = StateID(util.NewID("state"))
	next.Version = current.Version + 1
	next.Model.Prompt = input
	next.Model.Response = output

	transition := Transition{
		ID:          TransitionID(util.NewID("trans")),
		FromStateID: current.ID,
		ToStateID:   next.ID,
		Action:      action,
		Input:       input,
		Output:      output,
	}

	return next, transition
}
