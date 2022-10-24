package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/gustavo-bordin/x9/x9/repository"
)

type Resolver struct {
	repo repository.Repository
}

func NewResolver(repo repository.Repository) Resolver {
	resolver := Resolver{
		repo: repo,
	}

	return resolver
}

func (r Resolver) urls(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r Resolver) urlById(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r Resolver) addUrl(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}

func (r Resolver) deleteUrl(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
