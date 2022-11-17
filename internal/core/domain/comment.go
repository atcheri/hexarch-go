package domain

import (
	"time"
)

// Comment is a struct that
type Comment struct {
	userID    string
	key       string
	text      string
	createdAt time.Time
}
