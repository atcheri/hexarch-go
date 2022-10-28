package domain

// LanguageTranslation is
type LanguageTranslation struct {
	id   string
	code string
	text string
}

func NewLanguageTranslation(code, text string) (LanguageTranslation, error) {
	// validate code and text here
	return LanguageTranslation{
		code: code,
		text: text,
	}, nil
}

func (t LanguageTranslation) GetCode() string {
	return t.code
}
func (t LanguageTranslation) GetText() string {
	return t.text
}
