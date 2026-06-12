package runtime

type Model interface {
	Generate(prompt string, ctx map[string]any) (string, string, error)
}

type StaticModel struct {
	Response  string
	Reasoning string
}

func (m StaticModel) Generate(_ string, _ map[string]any) (string, string, error) {
	return m.Response, m.Reasoning, nil
}
