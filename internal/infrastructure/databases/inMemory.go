package databases

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

var (
	inMemoryProjects  = []string{"acme-test", "hitgub"}
	inMemoryLanguages = []string{"en", "fr", "ja", "pt", "es", "it"}
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

func (db *InMemoryDB) CreateProject(_ context.Context, name string) error {
	_, exists := lo.Find[string](db.projects, func(p string) bool {
		return p == name
	})
	if exists {
		return fmt.Errorf("cannot create the project %s. Project of the same name already exists", name)
	}

	db.projects = append(db.projects, name)

	return nil
}

func (db *InMemoryDB) EditProject(_ context.Context, oldName, newName string) error {
	_, i, _ := lo.FindIndexOf[string](db.projects, func(p string) bool {
		return p == oldName
	})
	if i == -1 {
		return fmt.Errorf("cannot edit the project %s. The project doesn't exist", oldName)
	}

	db.projects[i] = newName

	return nil
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
	translations, err := db.findProjectTranslationsForKey(name, key)
	if err != nil {
		return err
	}

	newTranslation, _ := domain.NewTranslation(name, key)
	for _, l := range inMemoryLanguages {
		textToSave := ""
		if l == code {
			textToSave = text
		}
		newTranslation, _ = newTranslation.AddTranslationWithCodeAndText(l, textToSave)
	}

	db.translations[name] = append(translations, newTranslation)

	return nil
}

// EditProjectTranslation just edits a translation for a given language, key and project
func (db *InMemoryDB) EditProjectTranslation(_ context.Context, id, key, code, text string) error {
	name := ""
	for projectName, translations := range db.translations {
		for _, translation := range translations {
			if translation.GetKey() == key {
				name = projectName
			}
		}
	}

	if name == "" {
		return fmt.Errorf("impossible to edit translation for the key %s. The key doesn't belong to any project", key)
	}

	translations, err := db.findProjectTranslations(name)
	if err != nil {
		return err
	}

	translation, ti, tFound := lo.FindIndexOf[domain.Translation](translations, func(t domain.Translation) bool {
		return t.GetKey() == key
	})

	if !tFound {
		return fmt.Errorf("impossible to edit translation for the key %s. The key translation key doesn't exist for this project %s", key, name)
	}

	languageTranslations := translation.GetTranslations()

	newTranslationLanguages := make([]domain.LanguageTranslation, len(languageTranslations))
	copy(newTranslationLanguages, languageTranslations)
	_, lti, ltFound := lo.FindIndexOf[domain.LanguageTranslation](languageTranslations, func(lt domain.LanguageTranslation) bool {
		return lt.GetCode() == code
	})

	if !ltFound {
		return fmt.Errorf("impossible to edit translation for the key %s and code %s. The key translation for the given language doesn't exist for this project %s", key, code, name)
	}

	newTranslationLanguage, _ := domain.NewLanguageTranslation(code, text)
	newTranslationLanguages[lti] = newTranslationLanguage

	newTranslations := make([]domain.Translation, len(translations))
	copy(newTranslations, translations)
	newTranslation, _ := domain.NewTranslation(name, key)
	lo.ForEach[domain.LanguageTranslation](newTranslationLanguages, func(tl domain.LanguageTranslation, index int) {
		newTranslation = newTranslation.AddTranslation(tl)
	})
	newTranslations[ti] = newTranslation
	db.translations[name] = newTranslations

	return nil
}

func (db *InMemoryDB) DeleteByKey(_ context.Context, name, key string) error {
	translations, err := db.findProjectTranslations(name)
	if err != nil {
		return err
	}

	_, _, tFound := lo.FindIndexOf[domain.Translation](translations, func(t domain.Translation) bool {
		return t.GetKey() == key
	})

	if !tFound {
		return fmt.Errorf("impossible to edit translation for the key %s. The key translation key doesn't exist for this project %s", key, name)
	}

	db.translations[name] = lo.Filter[domain.Translation](translations, func(t domain.Translation, _ int) bool {
		return t.GetKey() != key
	})

	return nil

}

func (db *InMemoryDB) findProjectTranslationsForKey(name, key string) ([]domain.Translation, error) {
	translations, err := db.findProjectTranslations(name)
	if err != nil {
		return nil, err
	}

	_, hasKey := lo.Find[domain.Translation](translations, func(t domain.Translation) bool {
		return t.GetKey() == key
	})

	if hasKey {
		return nil, fmt.Errorf("cannot add a new translation. The key %s already exists for this project %s", key, name)
	}
	return translations, nil
}

func (db *InMemoryDB) findProjectTranslations(name string) ([]domain.Translation, error) {
	translations, ok := db.translations[name]
	if !ok {
		return nil, fmt.Errorf("impossible to find translations for this project: %s", name)
	}

	return translations, nil
}

func createProjectTranslations(names, keys []string, languages []string, translationCodesAndValues [][]string) projectTranslationsType {
	var t projectTranslationsType = make(map[string][]domain.Translation)
	for _, name := range names {
		translations := make([]domain.Translation, len(keys))
		for ki, key := range keys {
			translation, _ := domain.NewTranslation(name, key)
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
