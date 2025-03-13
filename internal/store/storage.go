package store

import (
	"database/sql"
	"context"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, *Post) error
	}
	Users interface {
		Create(ctx context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UsersStore{db: db},
	}
}
