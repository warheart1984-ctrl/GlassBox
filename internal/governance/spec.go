package governance

import "glass-box/internal/constraints"

type Policy struct {
	ID          string                   `json:"id"`
	Version     string                   `json:"version"`
	Constraints []constraints.Constraint `json:"constraints"`
}

func ValidatePolicy(policy Policy) []error {
	errs := []error{}
	if policy.ID == "" {
		errs = append(errs, ValidationError{Field: "id", Message: "missing policy id"})
	}
	if policy.Version == "" {
		errs = append(errs, ValidationError{Field: "version", Message: "missing policy version"})
	}
	for _, c := range policy.Constraints {
		if c.Name == "" {
			errs = append(errs, ValidationError{Field: "constraint.name", Message: "missing constraint name"})
		}
		if c.Raw == "" {
			errs = append(errs, ValidationError{Field: "constraint.raw", Message: "missing constraint expression"})
		}
	}
	return errs
}
