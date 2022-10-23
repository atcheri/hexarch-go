package databases

import (
	"context"
	"fmt"

	"golang.org/x/exp/maps"

	"github.com/atcheri/hexarch-go/internal/core/domain"
	"github.com/pkg/errors"
)

type projectTranslationsType map[string][]domain.Translation

type InMemoryDB struct {
	words, sentences map[string]string
	translations     projectTranslationsType
}

// NewInMemoryDB is the factory function for a InMemoryDB struct
func NewInMemoryDB() *InMemoryDB {
	words := make(map[string]string, 0)
	words["firstName"] = "Prénom"
	words["middle_name"] = "Deuxième prénom"
	words["lastName"] = "Nom de famille"
	words["gender"] = "Sexe"
	words["birthday"] = "Date de naissance"
	words["title"] = "Titre"
	words["height"] = "Taille"
	sentences := make(map[string]string, 0)

	translations := createProjectTranslations([]string{"acme-test", "hitgub"}, []string{"home", "contact", "about-us"}, []string{"en", "fr", "pt", "jp"}, [][]string{
		{"home", "accueil", "casa", "ホーム"},
		{"contact", "contact", "contato", "問い合わせ"},
		{"about us", "à propos", "sobre nós", "会社概要"},
	})

	return &InMemoryDB{
		words:        words,
		sentences:    sentences,
		translations: translations,
	}
}

func (db *InMemoryDB) GetWords(offset, limit int) map[string]string {
	words := make(map[string]string)
	i := 0
	for key, word := range db.words {
		if i < offset {
			i++
			continue
		}
		if i == offset+limit {
			break
		}

		words[key] = word
		i++
	}
	return words
}

func (db *InMemoryDB) GetWordsInString(offset, limit int) []string {
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

func (db *InMemoryDB) GetProjectTranslations(ctx context.Context, name string, offset, limit int) ([]domain.Translation, error) {
	translations, ok := db.translations[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("no translations found for this project: %s", name))
	}

	//words := make(map[string]string)
	//i := 0
	//for key, word := range db.words {
	//	if i < offset {
	//		i++
	//		continue
	//	}
	//	if i == offset+limit {
	//		break
	//	}
	//
	//	words[key] = word
	//	i++
	//}
	return translations, nil
}

func createProjectTranslations(names, keys []string, languages []string, translationCodesAndValues [][]string) projectTranslationsType {
	var t projectTranslationsType = make(map[string][]domain.Translation)
	for _, name := range names {
		translations := make([]domain.Translation, len(keys))
		for ki, key := range keys {
			translation, _ := domain.NewTranslation(key)
			for li, lang := range languages {
				tv, _ := domain.NewTranslationValue(lang, translationCodesAndValues[ki][li])
				translation = translation.AddTranslation(tv)
			}
			translations[ki] = translation
		}
		t[name] = translations
	}

	return t
}
