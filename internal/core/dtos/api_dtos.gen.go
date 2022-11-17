// Package dto provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package dto

import (
	"time"
)

// CommentDTO defines model for CommentDTO.
type CommentDTO struct {
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
	UserId    string    `json:"userId"`
}

// The payload body to create a new project
type CreateProjectRequestBody struct {
	Name string `json:"name"`
}

// The payload body to create a translation
type CreateProjectTranslationRequestBody struct {
	Code string `json:"code"`
	Key  string `json:"key"`
	Text string `json:"text"`
}

// The payload body to create a comment for a translation key
type CreateTranslationCommentRequestBody struct {
	Text   string `json:"text"`
	UserId string `json:"userId"`
}

// The payload body to delete translations for a key
type DeleteProjectTranslationRequestBody struct {
	Key string `json:"key"`
}

// EditProjectRequestBody defines model for EditProjectRequestBody.
type EditProjectRequestBody struct {
	// Embedded struct due to allOf(#/components/schemas/CreateProjectRequestBody)
	CreateProjectRequestBody `yaml:",inline"`
}

// EditProjectTranslationRequestBody defines model for EditProjectTranslationRequestBody.
type EditProjectTranslationRequestBody struct {
	// Embedded struct due to allOf(#/components/schemas/CreateProjectTranslationRequestBody)
	CreateProjectTranslationRequestBody `yaml:",inline"`
}

// Standard Error response.
type Error struct {
	ErrorType string `json:"error_type"`

	// The detailed reason of the error.
	Message string `json:"message"`

	// Error name.
	Name string `json:"name"`
}

// NotFoundError defines model for NotFoundError.
type NotFoundError struct {
	// Embedded struct due to allOf(#/components/schemas/Error)
	Error `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// example `The translations not found for this project`
	ErrorType string `json:"error_type"`
}

// TranslationDTO defines model for TranslationDTO.
type TranslationDTO struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

// TranslationKeyDTO defines model for TranslationKeyDTO.
type TranslationKeyDTO struct {
	Id        string           `json:"id"`
	Key       string           `json:"key"`
	Languages []TranslationDTO `json:"languages"`
}

// GetProjectTranslationsLanguagesParam defines model for GetProjectTranslationsLanguagesParam.
type GetProjectTranslationsLanguagesParam []string

// ProjectNameParam defines model for ProjectNameParam.
type ProjectNameParam string

// TranslationIdParam defines model for TranslationIdParam.
type TranslationIdParam string

// Standard Error response.
type BadRequest Error

// Standard Error response.
type InternalServerError Error

// Ressource not found Error response.
type NotFound NotFoundError

// ProjectTranslationsGetResponse defines model for ProjectTranslationsGetResponse.
type ProjectTranslationsGetResponse struct {
	// The total amount of translations
	Total int `json:"total"`

	// The list of translations
	Translations []TranslationKeyDTO `json:"translations"`
}

// TranslationCommentsGetResponse defines model for TranslationCommentsGetResponse.
type TranslationCommentsGetResponse struct {
	// The total amount of comments for the translation key
	Total int `json:"total"`

	// The list of comments
	Translations *[]CommentDTO `json:"translations,omitempty"`
}

// PostProjectJSONBody defines parameters for PostProject.
type PostProjectJSONBody CreateProjectRequestBody

// PutProjectJSONBody defines parameters for PutProject.
type PutProjectJSONBody EditProjectRequestBody

// PostCommentJSONBody defines parameters for PostComment.
type PostCommentJSONBody CreateTranslationCommentRequestBody

// DeleteProjectTranslationsJSONBody defines parameters for DeleteProjectTranslations.
type DeleteProjectTranslationsJSONBody DeleteProjectTranslationRequestBody

// GetProjectTranslationsParams defines parameters for GetProjectTranslations.
type GetProjectTranslationsParams struct {
	// A list of Language codes
	Languages *GetProjectTranslationsLanguagesParam `json:"languages,omitempty"`
}

// PostProjectTranslationJSONBody defines parameters for PostProjectTranslation.
type PostProjectTranslationJSONBody CreateProjectTranslationRequestBody

// PutProjectTranslationJSONBody defines parameters for PutProjectTranslation.
type PutProjectTranslationJSONBody EditProjectTranslationRequestBody

// PostProjectJSONRequestBody defines body for PostProject for application/json ContentType.
type PostProjectJSONRequestBody PostProjectJSONBody

// PutProjectJSONRequestBody defines body for PutProject for application/json ContentType.
type PutProjectJSONRequestBody PutProjectJSONBody

// PostCommentJSONRequestBody defines body for PostComment for application/json ContentType.
type PostCommentJSONRequestBody PostCommentJSONBody

// DeleteProjectTranslationsJSONRequestBody defines body for DeleteProjectTranslations for application/json ContentType.
type DeleteProjectTranslationsJSONRequestBody DeleteProjectTranslationsJSONBody

// PostProjectTranslationJSONRequestBody defines body for PostProjectTranslation for application/json ContentType.
type PostProjectTranslationJSONRequestBody PostProjectTranslationJSONBody

// PutProjectTranslationJSONRequestBody defines body for PutProjectTranslation for application/json ContentType.
type PutProjectTranslationJSONRequestBody PutProjectTranslationJSONBody
