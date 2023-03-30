package constant

type AIMode int

const (
	GPT AIMode = iota
	DELL
)

var AIModeNames = map[AIMode]string{
	GPT:  "GPT",
	DELL: "DELL",
}

func IsAIMode(mode AIMode) bool {
	_, ok := AIModeNames[mode]
	return ok
}
