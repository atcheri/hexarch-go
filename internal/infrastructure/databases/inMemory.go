package databases

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v4"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/samber/lo"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

var (
	inMemoryProjects  = []string{"acme-test", "hitgub"}
	inMemoryLanguages = []string{"en", "fr", "ja", "pt", "es", "it"}
)

type projectTranslationsType map[string][]domain.Translation

// InMemoryDB as it's name indicates, is a in memory database
type InMemoryDB struct {
	projects     []string
	translations projectTranslationsType
	comments     []domain.Comment
}

// NewInMemoryDB is the factory function for a InMemoryDB struct
func NewInMemoryDB() *InMemoryDB {
	projectTranslations := createProjectTranslations(inMemoryProjects, []string{"home", "contact", "about-us"}, []string{"en", "fr", "pt", "jp"}, [][]string{
		{"home", "accueil", "casa", "ホーム"},
		{"contact", "contact", "contato", "問い合わせ"},
		{"about us", "à propos", "sobre nós", "会社概要"},
	})

	ids := make([]string, 0)
	for _, translations := range projectTranslations {
		for _, translation := range translations {
			ids = append(ids, translation.GetID())
		}
	}

	comments := createTranslationComments(ids)

	return &InMemoryDB{
		projects:     inMemoryProjects,
		translations: projectTranslations,
		comments:     comments,
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

// GetProjectTranslation gets a specific translation in a project
func (db *InMemoryDB) GetProjectTranslation(_ context.Context, id string) (domain.Translation, error) {
	name, err := db.findProjectNameFromPredicate(id, func(t domain.Translation) bool {
		return t.GetID() == id
	})

	if err != nil {
		return domain.Translation{}, err
	}

	translations, err := db.findProjectTranslations(name)
	if err != nil {
		return domain.Translation{}, err
	}

	translation, ok := lo.Find[domain.Translation](translations, func(t domain.Translation) bool {
		return t.GetID() == id
	})
	if !ok {
		return domain.Translation{}, fmt.Errorf("couldn't find the translation for the id %s", id)
	}

	return translation, nil
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

// AddProjectTranslationComment adds a comment to an existing translation
func (db *InMemoryDB) AddProjectTranslationComment(_ context.Context, id, userID, text string) error {
	db.comments = append(db.comments, domain.NewComment(userID, id, text))

	return nil
}

// EditProjectTranslation just edits a translation for a given language, key and project
func (db *InMemoryDB) EditProjectTranslation(_ context.Context, _, key, code, text string) error {
	name, err := db.findProjectNameFromPredicate(key, func(t domain.Translation) bool {
		return t.GetKey() == key
	})
	if err != nil {
		return err
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

// GetTranslationComments filters and only keeps comments that belong to a specific translation-id
func (db *InMemoryDB) GetTranslationComments(_ context.Context, id string) ([]domain.Comment, error) {
	return lo.Filter[domain.Comment](db.comments, func(c domain.Comment, _ int) bool {
		return c.GetTranslationID() == id
	}), nil
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

func (db *InMemoryDB) findProjectNameFromPredicate(param string, predicate func(t domain.Translation) bool) (string, error) {
	name := ""
	for projectName, translations := range db.translations {
		for _, translation := range translations {
			if predicate(translation) {
				name = projectName
			}
		}
	}

	if name == "" {
		return "", fmt.Errorf("impossible to edit translation for the param %s. The key doesn't belong to any project", param)
	}

	return name, nil
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

func createTranslationComments(ids []string) []domain.Comment {
	comments := make([]domain.Comment, 0)
	for _, id := range ids {
		randomCount := rand.Intn(5) + 1
		for i := 0; i < randomCount; i++ {
			uniqID, _ := uuid.NewV4()
			comments = append(comments, domain.NewComment(uniqID.String(), id, faker.Sentence()))
		}
	}

	return comments
}
