package storage

import (
	"context"
	"time"
)

type Note struct {
	ID        int64     `db:"note_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type Storage interface {
	Save(ctx context.Context, note Note) (*Note, error)
	UserNotes(ctx context.Context, id int64) ([]*Note, error)
	PickRandom(ctx context.Context, id int64) (*Note, error)
	Remove(ctx context.Context, note Note) error
}
