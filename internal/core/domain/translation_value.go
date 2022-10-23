package domain

// TranslationValue is
type TranslationValue struct {
	code string
	text string
}

func NewTranslationValue(code, text string) (TranslationValue, error) {
	// validate code and text here
	return TranslationValue{
		code: code,
		text: text,
	}, nil
}

func (t TranslationValue) GetCode() string {
	return t.code
}
func (t TranslationValue) GetText() string {
	return t.text
}
