package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	// Defining all necessary methods in an interface
	PostStoreAllMethods interface {
		GetAll() ([]*Post, error) // Add this method
		// other methods...
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UsersStore{db: db},
	}
}
