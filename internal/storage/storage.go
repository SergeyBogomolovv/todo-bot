package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SergeyBogomolovv/notes-bot/pkg/e"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *storage {
	return &storage{db: db}
}

func (s *storage) Save(ctx context.Context, note Note) (*Note, error) {
	query := `
  INSERT INTO notes (title, content, user_id)
  VALUES ($1, $2, $3)
  RETURNING note_id, title, content, user_id, created_at`
	if err := s.db.GetContext(ctx, &note, query, note.Title, note.Content, note.UserID); err != nil {
		return nil, e.Wrap("can't save note", err)
	}
	return &note, nil
}

func (s *storage) UserNotes(ctx context.Context, id int64) ([]*Note, error) {
	query := `
  SELECT note_id, title, content, user_id, created_at
  FROM notes
  WHERE user_id = $1
  ORDER BY created_at`
	var res []*Note
	if err := s.db.SelectContext(ctx, &res, query, id); err != nil {
		return nil, e.Wrap("can't get notes", err)
	}
	return res, nil
}

func (s *storage) PickRandom(ctx context.Context, id int64) (*Note, error) {
	query := `
  SELECT note_id, title, content, user_id, created_at
  FROM notes
  WHERE user_id = $1
  ORDER BY RANDOM()
  LIMIT 1`

	note := new(Note)
	if err := s.db.GetContext(ctx, note, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoNotes
		}
		return nil, e.Wrap("can't pick random note", err)
	}
	return note, nil
}

func (s *storage) Remove(ctx context.Context, note Note) error {
	query := `DELETE FROM notes WHERE note_id = $1`
	res, err := s.db.ExecContext(ctx, query, note.ID)
	if err != nil {
		return e.Wrap("can't remove note", err)
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return e.Wrap("can't check affected rows", err)
	}
	if aff == 0 {
		return ErrNoteNotFound
	}
	return nil
}
