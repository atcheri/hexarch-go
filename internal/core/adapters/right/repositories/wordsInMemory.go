package adapters

import (
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

type inMemoryWords struct {
	db *databases.InMemoryDB
}

// NewInMemoryWords instantiates a new inMemorySentences that implements WordsRepository interface
func NewInMemoryWords(db *databases.InMemoryDB) ports.WordsRepository {
	return inMemoryWords{db: db}
}

func (i inMemoryWords) GetAll(offset, limit int) []string {
	return i.db.GetWords(offset, limit)
}

func (i inMemoryWords) GetByKey(key string) string {
	//TODO implement me
	panic("implement me")
}

func (i inMemoryWords) SetWord(key, content string) {
	i.db.AddWord(key, content)
}
