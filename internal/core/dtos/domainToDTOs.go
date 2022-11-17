package dto

import (
	"github.com/samber/lo"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

func ToTranslationKeyDTOs(ts []domain.Translation) []TranslationKeyDTO {
	return lo.Map[domain.Translation, TranslationKeyDTO](ts, func(t domain.Translation, _ int) TranslationKeyDTO {
		return ToTranslationKeyDTO(t)
	})
}

func ToTranslationKeyDTO(t domain.Translation) TranslationKeyDTO {
	languages := t.GetTranslations()
	languageDTOs := make([]TranslationDTO, len(languages))
	for i, l := range languages {
		languageDTOs[i] = TranslationDTO{
			Code: l.GetCode(),
			Text: l.GetText(),
		}
	}
	return TranslationKeyDTO{
		Id:        t.GetID(),
		Key:       t.GetKey(),
		Languages: languageDTOs,
	}
}
