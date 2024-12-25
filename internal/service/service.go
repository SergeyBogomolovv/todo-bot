package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/SergeyBogomolovv/notes-bot/internal/storage"
)

type service struct {
	log     *slog.Logger
	storage storage.Storage
}

func New(log *slog.Logger, storage storage.Storage) *service {
	return &service{log: log, storage: storage}
}

func (s *service) UserNotesTitles(ctx context.Context, id int64) ([]string, error) {
	notes, err := s.storage.UserNotes(ctx, id)
	if err != nil {
		s.log.Error("can't get user notes", "error", err)
		return nil, err
	}
	titles := make([]string, len(notes))
	for i, note := range notes {
		titles[i] = note.Title
	}
	return titles, nil
}

func (s *service) CreateNote(ctx context.Context, title, content string, id int64) (*NoteResponse, error) {
	payload := storage.Note{
		Title:   title,
		Content: content,
		UserID:  id,
	}
	note, err := s.storage.Save(ctx, payload)
	if err != nil {
		s.log.Error("can't create note", "error", err)
		return nil, err
	}
	res := &NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}
	return res, nil
}

func (s *service) PickRandomNote(ctx context.Context, userId int64) (*NoteResponse, error) {
	note, err := s.storage.PickRandom(ctx, userId)
	if err != nil {
		if errors.Is(err, storage.ErrNoNotes) {
			s.log.Info("no notes for user", "user_id", userId)
			return nil, ErrNoNotes
		}
		s.log.Error("can't pick random note", "error", err)
		return nil, err
	}
	res := &NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}
	return res, nil
}

func (s *service) RemoveNote(ctx context.Context, id int64) error {
	note := storage.Note{ID: id}
	err := s.storage.Remove(ctx, note)
	if err != nil {
		s.log.Error("can't remove note", "error", err)
		return err
	}
	return nil
}
