package runtime

import (
	"context"
	"errors"
	"time"

	"glass-box/internal/classifier"
	"glass-box/internal/constraints"
	"glass-box/internal/ledger"
	"glass-box/internal/state"
)

type Executor struct {
	Model         Model
	Classifiers   []classifier.Classifier
	Constraints   []constraints.Constraint
	Secret        []byte
	PolicyVersion string
}

func (e Executor) Execute(ctx context.Context, current state.State, input string) (state.State, Receipt, ledger.Record, error) {
	if e.Model == nil {
		return current, Receipt{}, ledger.Record{}, errors.New("runtime model is required")
	}
	if err := ctx.Err(); err != nil {
		return current, Receipt{}, ledger.Record{}, err
	}

	decisions := []classifier.Decision{}
	for _, cls := range e.Classifiers {
		out, err := cls.Evaluate(input, current.Context)
		if err != nil {
			return current, Receipt{}, ledger.Record{}, err
		}
		decisions = append(decisions, out...)
	}

	response, reasoning, err := e.Model.Generate(input, current.Context)
	if err != nil {
		return current, Receipt{}, ledger.Record{}, err
	}

	next, transition := state.ApplyTransition(current, "generate", input, response)
	next.Model.Reasoning = reasoning

	results := []constraints.ConstraintResult{}
	constraintCtx := map[string]any{
		"input":     input,
		"response":  response,
		"reasoning": reasoning,
		"action":    transition.Action,
	}
	for _, c := range e.Constraints {
		result := constraints.Evaluate(c, constraintCtx)
		results = append(results, result)
		if c.Severity == constraints.SeverityBlock && result.Status == constraints.ResultFail {
			return current, Receipt{}, ledger.Record{}, errors.New(result.Details)
		}
	}
	next.Constraints.Results = make([]state.ConstraintResult, 0, len(results))
	for _, result := range results {
		next.Constraints.Results = append(next.Constraints.Results, state.ConstraintResult{
			Name:    result.Name,
			Status:  string(result.Status),
			Details: result.Details,
		})
	}

	receipt := Receipt{
		Timestamp:           time.Now().UTC(),
		Input:               input,
		Response:            response,
		Reasoning:           reasoning,
		ClassifierDecisions: decisions,
		ConstraintResults:   results,
		PolicyVersion:       e.PolicyVersion,
	}

	record := ledger.NewRecord(ledger.RecordTransition, map[string]any{
		"transition": transition,
		"receipt":    receipt,
	})
	signed, err := ledger.SignRecord(record, e.Secret)
	if err != nil {
		return current, Receipt{}, ledger.Record{}, err
	}

	return next, receipt, signed, nil
}
