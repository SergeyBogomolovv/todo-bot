package controller

import (
	"context"

	"github.com/SergeyBogomolovv/notes-bot/internal/service"
)

type UserState struct {
	State State
	Data  any
}

type State int

const (
	WaitingForTitle   State = iota
	WaitingForContent State = iota
)

type NotesService interface {
	UserNotesTitles(ctx context.Context, id int64) ([]string, error)
	CreateNote(ctx context.Context, title, content string, userId int64) (*service.NoteResponse, error)
	PickRandomNote(ctx context.Context, userId int64) (*service.NoteResponse, error)
	RemoveNote(ctx context.Context, id int64) error
}
