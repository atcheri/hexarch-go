package domain

import "github.com/samber/lo"

// Project is
type Project struct {
	name         string
	translations []Translation
	languages    []LanguageCodeAndName
}

func (p Project) GetName() string {
	return p.name
}

func (p Project) GetLanguageCodes() []string {
	return getLanguageBy(p.languages, "code")
}

func (p Project) GetLanguageNames() []string {
	return getLanguageBy(p.languages, "name")
}

func (p Project) GetLanguageNaticeNames() []string {
	return getLanguageBy(p.languages, "native")
}

func getLanguageBy(languages []LanguageCodeAndName, fieldName string) []string {
	return lo.Map[LanguageCodeAndName, string](languages, func(l LanguageCodeAndName, index int) string {
		switch fieldName {
		case "code":
			return l.Code
		case "name":
			return l.Name
		case "native":
			return l.NativeName
		default:
			return ""
		}
	})
}
