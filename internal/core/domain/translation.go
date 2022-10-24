package domain

import (
	"fmt"

	"github.com/pkg/errors"
)

// Translation is a struct that contains the key and the different language translations
type Translation struct {
	key                  string
	languageTranslations []LanguageTranslation
}

func NewTranslation(key string) (Translation, error) {
	// validate key here
	return Translation{
		key:                  key,
		languageTranslations: make([]LanguageTranslation, 0),
	}, nil
}

func (t Translation) AddTranslation(l LanguageTranslation) Translation {
	return Translation{
		key:                  t.key,
		languageTranslations: append(t.languageTranslations, l),
	}
}

func (t Translation) GetKey() string {
	return t.key
}

func (t Translation) GetTranslations() []LanguageTranslation {
	return t.languageTranslations
}

func (t Translation) AddTranslationWithCodeAndText(code, text string) (Translation, error) {
	translation, err := NewLanguageTranslation(code, text)
	if err != nil {
		return Translation{}, errors.Wrap(err, fmt.Sprintf("cannot add translation with code %s and text %s", code, text))
	}

	return t.AddTranslation(translation), nil
}
