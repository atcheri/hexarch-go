package ports

type SentencesRepository interface {
	GetByKey(key string) string
	SetSentence(key string, content string)
}
