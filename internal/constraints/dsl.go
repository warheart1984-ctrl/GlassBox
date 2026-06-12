package constraints

import (
	"fmt"
	"strings"
)

type Severity string

const (
	SeverityBlock Severity = "block"
	SeverityWarn  Severity = "warn"
)

type ResultStatus string

const (
	ResultPass ResultStatus = "pass"
	ResultFail ResultStatus = "fail"
)

type Constraint struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Raw         string   `json:"raw"`
	Severity    Severity `json:"severity"`
}

type ConstraintResult struct {
	Name    string       `json:"name"`
	Status  ResultStatus `json:"status"`
	Details string       `json:"details"`
}

func Evaluate(c Constraint, ctx map[string]any) ConstraintResult {
	raw := strings.TrimSpace(c.Raw)
	if strings.HasPrefix(raw, "deny ") {
		matches, details := evalComparison(strings.TrimSpace(strings.TrimPrefix(raw, "deny ")), ctx)
		if matches {
			return ConstraintResult{Name: c.Name, Status: ResultFail, Details: details}
		}
		return ConstraintResult{Name: c.Name, Status: ResultPass, Details: "deny expression did not match"}
	}

	if strings.HasPrefix(raw, "require ") {
		matches, details := evalComparison(strings.TrimSpace(strings.TrimPrefix(raw, "require ")), ctx)
		if matches {
			return ConstraintResult{Name: c.Name, Status: ResultPass, Details: details}
		}
		return ConstraintResult{Name: c.Name, Status: ResultFail, Details: details}
	}

	return ConstraintResult{Name: c.Name, Status: ResultFail, Details: "unsupported constraint expression"}
}

func evalComparison(expr string, ctx map[string]any) (bool, string) {
	var op string
	switch {
	case strings.Contains(expr, "!="):
		op = "!="
	case strings.Contains(expr, "=="):
		op = "=="
	default:
		return false, fmt.Sprintf("unsupported comparison: %s", expr)
	}

	parts := strings.SplitN(expr, op, 2)
	if len(parts) != 2 {
		return false, fmt.Sprintf("invalid comparison: %s", expr)
	}

	key := strings.TrimSpace(parts[0])
	want := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
	got, ok := ctx[key]
	if !ok {
		got = ""
	}
	gotString := fmt.Sprint(got)

	switch op {
	case "==":
		return gotString == want, fmt.Sprintf("%s == %q evaluated with %q", key, want, gotString)
	case "!=":
		return gotString != want, fmt.Sprintf("%s != %q evaluated with %q", key, want, gotString)
	default:
		return false, fmt.Sprintf("unsupported operator: %s", op)
	}
}
