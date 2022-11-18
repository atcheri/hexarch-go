package domain

import (
	"time"
)

// Comment is a struct that
type Comment struct {
	userID        string
	translationID string
	text          string
	createdAt     time.Time
}

// NewComment is the Comment factory function
func NewComment(userID, translationID, text string) Comment {
	return Comment{
		userID:        userID,
		translationID: translationID,
		text:          text,
		createdAt:     time.Now(),
	}
}

// GetUserID returns the userID field
func (c Comment) GetUserID() string {
	return c.userID
}

// GetTranslationID returns the translationID field
func (c Comment) GetTranslationID() string {
	return c.translationID
}

// GetText returns the text field
func (c Comment) GetText() string {
	return c.text
}

// GetCreatedAt returns the createdAt field
func (c Comment) GetCreatedAt() time.Time {
	return c.createdAt
}
