package services

type AI interface {
	Setup()
	Reply(string) string
}
