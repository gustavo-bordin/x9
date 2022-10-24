package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Api struct {
	Port     string
	Handlers Handlers
}

func NewApi(port string, handlers Handlers) Api {
	return Api{
		Port:     port,
		Handlers: handlers,
	}
}

func (a Api) Start() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/urls", a.Handlers.Url)

	err := e.Start(a.Port)
	return err
}
