package api

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

type Handlers struct {
	urlSchema graphql.Schema
}

func NewHandlers(urlSchema graphql.Schema) Handlers {
	return Handlers{
		urlSchema: urlSchema,
	}
}

func (h Handlers) Url(c echo.Context) error {
	var params Params

	if err := c.Bind(&params); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	result := graphql.Do(graphql.Params{
		Context:        c.Request().Context(),
		Schema:         h.urlSchema,
		RequestString:  params.Query,
		VariableValues: params.Variables,
		OperationName:  params.Operation,
	})

	if len(result.Errors) > 0 {
		return c.String(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, result)
}
