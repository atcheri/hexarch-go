package databases

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type InMemoryDB struct {
	words, sentences map[string]string
}

// NewInMemoryDB is the factory function for a InMemoryDB struct
func NewInMemoryDB() *InMemoryDB {
	words := make(map[string]string, 0)
	words["firstName"] = "Prénom"
	words["middle_name"] = "Deuxième prénom"
	words["lastName"] = "Nom de famille"
	words["gender"] = "Sexe"
	words["bitrhday"] = "Date de naissance"
	words["title"] = "Titre"
	words["height"] = "Taille"
	sentences := make(map[string]string, 0)
	return &InMemoryDB{
		words:     words,
		sentences: sentences,
	}
}

func (db *InMemoryDB) GetWords(offset, limit int) []string {
	return maps.Values(db.words)[offset:limit]
}

func (db *InMemoryDB) GetWordByKey(key string) (string, error) {
	if w, ok := db.words[key]; ok {
		return w, nil
	}

	return "", fmt.Errorf("word not for for key %s", key)
}

func (db *InMemoryDB) AddWord(key, content string) {
	db.words[key] = content
}

func (db *InMemoryDB) RemoveWord(key string) {
	delete(db.words, key)
}
