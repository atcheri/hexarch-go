// Package dto exposes helper functions to convert domain models to dtos
package dto

import (
	"github.com/samber/lo"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

// ToTranslationKeyDTOs converts the domain model to it's corresponding DTO
func ToTranslationKeyDTOs(ts []domain.Translation) []TranslationKeyDTO {
	return lo.Map[domain.Translation, TranslationKeyDTO](ts, func(t domain.Translation, _ int) TranslationKeyDTO {
		return ToTranslationKeyDTO(t)
	})
}

// ToTranslationKeyDTO converts the domain model to it's corresponding DTO
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

// ToCommentDTO converts the domain model to it's corresponding DTO
func ToCommentDTO(c domain.Comment) CommentDTO {
	return CommentDTO{
		CreatedAt: c.GetCreatedAt(),
		Text:      c.GetText(),
		UserId:    c.GetUserID(),
	}
}
