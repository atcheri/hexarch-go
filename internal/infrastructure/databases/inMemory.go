package databases

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

var (
	inMemoryProjects  = []string{"acme-test", "hitgub"}
	inMemorylanguages = []string{"en", "fr", "ja", "pt", "es", "it"}
)

type projectTranslationsType map[string][]domain.Translation

type InMemoryDB struct {
	projects     []string
	translations projectTranslationsType
}

// NewInMemoryDB is the factory function for a InMemoryDB struct
func NewInMemoryDB() *InMemoryDB {
	translations := createProjectTranslations(inMemoryProjects, []string{"home", "contact", "about-us"}, []string{"en", "fr", "pt", "jp"}, [][]string{
		{"home", "accueil", "casa", "ホーム"},
		{"contact", "contact", "contato", "問い合わせ"},
		{"about us", "à propos", "sobre nós", "会社概要"},
	})

	return &InMemoryDB{
		projects:     inMemoryProjects,
		translations: translations,
	}
}

func (db *InMemoryDB) GetProjectTranslations(_ context.Context, name string, offset, limit int) ([]domain.Translation, error) {
	_, ok := db.translations[name]
	if !ok {
		return nil, fmt.Errorf("no translations found for this project: %s", name)
	}

	translations := make([]domain.Translation, 0)
	i := 0
	for _, t := range db.translations[name] {
		if i < offset {
			i++
			continue
		}
		if i == offset+limit {
			break
		}

		translations = append(translations, t)
		i++
	}

	return translations, nil
}

func (db *InMemoryDB) AddProjectTranslation(_ context.Context, name, key, code, text string) error {
	translations, ok := db.translations[name]
	if !ok {
		return fmt.Errorf("impossible to add translation to this project: %s", name)
	}

	_, hasKey := lo.Find[domain.Translation](translations, func(t domain.Translation) bool {
		return t.GetKey() == key
	})

	if hasKey {
		return fmt.Errorf("cannot add a new translation. The key %s already exists for this project %s", key, name)
	}

	newTranslation, _ := domain.NewTranslation(key)
	for _, l := range inMemorylanguages {
		textToSave := ""
		if l == code {
			textToSave = text
		}
		newTranslation, _ = newTranslation.AddTranslationWithCodeAndText(l, textToSave)
	}

	db.translations[name] = append(translations, newTranslation)

	return nil
}

func createProjectTranslations(names, keys []string, languages []string, translationCodesAndValues [][]string) projectTranslationsType {
	var t projectTranslationsType = make(map[string][]domain.Translation)
	for _, name := range names {
		translations := make([]domain.Translation, len(keys))
		for ki, key := range keys {
			translation, _ := domain.NewTranslation(key)
			for li, lang := range languages {
				tv, _ := domain.NewLanguageTranslation(lang, translationCodesAndValues[ki][li])
				translation = translation.AddTranslation(tv)
			}
			translations[ki] = translation
		}
		t[name] = translations
	}

	return t
}
