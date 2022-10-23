package domain

// Translation is a struct that contains the key and the different language translations
type Translation struct {
	key       string
	languages []TranslationValue
}

func NewTranslation(key string) (Translation, error) {
	// validate key here
	return Translation{
		key:       key,
		languages: make([]TranslationValue, 0),
	}, nil
}

func (t Translation) AddTranslation(l TranslationValue) Translation {
	return Translation{
		key:       t.key,
		languages: append(t.languages, l),
	}
}

func (t Translation) GetKey() string {
	return t.key
}

func (t Translation) GetTranslations() []TranslationValue {
	return t.languages
}
