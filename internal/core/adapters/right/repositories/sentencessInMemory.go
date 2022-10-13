package adapters

import (
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

type inMemorySentences struct {
	db *databases.InMemoryDB
}

// NewInMemorySentences instantiates a new inMemorySentences that implements SentencesRepository interface
func NewInMemorySentences(db *databases.InMemoryDB) ports.SentencesRepository {
	return inMemorySentences{db: db}
}

func (i inMemorySentences) GetByKey(key string) string {
	//TODO implement me
	panic("implement me")
}

func (i inMemorySentences) SetSentence(key string, content string) {
	//TODO implement me
	panic("implement me")
}
