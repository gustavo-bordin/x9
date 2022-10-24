package x9

import (
	"github.com/gustavo-bordin/x9/x9/api"
	graphqlApi "github.com/gustavo-bordin/x9/x9/api/graphql"
	"github.com/gustavo-bordin/x9/x9/config"
	"github.com/gustavo-bordin/x9/x9/repository"
)

func NewApplication(cfg config.Config) {
	db, err := repository.Connect(cfg)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	resolver := graphqlApi.NewResolver(repo)
	query := graphqlApi.NewQuery(resolver)

	mutations := query.GetUrlMutations()
	queries := query.GetUrlQueries()
	urlSchema, err := graphqlApi.NewUrlSchema(queries, mutations)

	handlers := api.NewHandlers(urlSchema)
	api := api.NewApi(":5005", handlers)

	err = api.Start()
	if err != nil {
		panic(err)
	}
}
