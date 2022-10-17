package x9

import (
	"database/sql"
	"errors"

	"github.com/gustavo-bordin/x9/x9/config"
	"github.com/gustavo-bordin/x9/x9/handlers"
	"github.com/gustavo-bordin/x9/x9/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	errDbIsRequired = errors.New("DB is required")
)

type X9 struct {
	db *sql.DB
}

func NewApplication(cfg config.Config) {
	db, err := repository.Connect(cfg)
	if err != nil {
		panic(err)
	}

	x9, err := newX9(db)
	if err != nil {
		panic(err)
	}

	err = x9.startApi()
	if err != nil {
		panic(err)
	}
}

func newX9(db *sql.DB) (X9, error) {
	if db == nil {
		return X9{}, errDbIsRequired
	}

	x9 := X9{db: db}
	return x9, nil
}

func (x X9) startApi() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/track", handlers.Track)

	// Start server
	err := e.Start(":5005")
	return err
}
