package service

import (
	"errors"
	"time"
)

type NoteResponse struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
}

var (
	ErrNoNotes      = errors.New("no notes")
	ErrNoteNotFound = errors.New("note not found")
)
