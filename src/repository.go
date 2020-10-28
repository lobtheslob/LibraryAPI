package main

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
)

//NewRepository connects to the sql db
func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

type Repository struct {
	db *sql.DB
}

func (r Repository) store(ctx context.Context, book Book) error {
	query := `INSERT INTO book (id, name, author, isbn) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE id = ?, name = ?, author = ?, isbn = ?`
	log.Info("book in repo: ", book)
	_, err := r.db.ExecContext(
		ctx,
		query,
		book.ID,
		book.Name,
		book.Author,
		book.ISBN,
		book.ID, // start of upsert
		book.Name,
		book.Author,
		book.ISBN,
	)
	if err != nil {
		log.Error("Error while storing model")
		return err
	}

	return nil
}