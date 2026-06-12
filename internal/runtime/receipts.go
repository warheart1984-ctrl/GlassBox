package runtime

import (
	"time"

	"glass-box/internal/classifier"
	"glass-box/internal/constraints"
)

type Receipt struct {
	Timestamp           time.Time                      `json:"timestamp"`
	Input               string                         `json:"input"`
	Response            string                         `json:"response"`
	Reasoning           string                         `json:"reasoning"`
	ClassifierDecisions []classifier.Decision          `json:"classifier_decisions"`
	ConstraintResults   []constraints.ConstraintResult `json:"constraint_results"`
	PolicyVersion       string                         `json:"policy_version"`
}
