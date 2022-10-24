package repository

import (
	"database/sql"
	"fmt"

	"github.com/gustavo-bordin/x9/x9/config"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func Connect(cfg config.Config) (*sql.DB, error) {
	connUri := createConnUri(cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME)
	db, err := sql.Open("postgres", connUri)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnUri(dbUser, dbPass, dbName string) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPass, dbName,
	)
}
