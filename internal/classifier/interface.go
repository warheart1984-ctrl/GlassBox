package classifier

type Decision struct {
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Confidence float64 `json:"confidence"`
	Reason     string  `json:"reason"`
	Raw        any     `json:"raw"`
}

type Classifier interface {
	ID() string
	Evaluate(input string, ctx map[string]any) ([]Decision, error)
}

type StaticClassifier struct {
	ClassifierID string
	Decisions    []Decision
}

func (c StaticClassifier) ID() string {
	return c.ClassifierID
}

func (c StaticClassifier) Evaluate(_ string, _ map[string]any) ([]Decision, error) {
	return c.Decisions, nil
}
