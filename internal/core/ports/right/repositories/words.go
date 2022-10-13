package ports

type WordsRepository interface {
	GetAll(offset, limit int) []string
	GetByKey(key string) string
	SetWord(key string, content string)
	RemoveWord(key string) error
}
