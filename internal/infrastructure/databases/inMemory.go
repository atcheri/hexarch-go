package databases

import "golang.org/x/exp/maps"

type InMemoryDB struct {
	words, sentences map[string]string
}

// NewInMemoryDB is the factory function for a InMemoryDB struct
func NewInMemoryDB() *InMemoryDB {
	words := make(map[string]string, 0)
	words["firstName"] = "Prénom"
	words["lastName"] = "Nom de famille"
	sentences := make(map[string]string, 0)
	return &InMemoryDB{
		words:     words,
		sentences: sentences,
	}
}

func (db *InMemoryDB) GetWords(offset, limit int) []string {
	return maps.Values(db.words)
}

func (db *InMemoryDB) GetWordByKey(key string) string {
	if w, ok := db.words[key]; ok {
		return w
	}

	return ""
}

func (db *InMemoryDB) AddWord(key, content string) {
	db.words[key] = content
}
