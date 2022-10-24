package graphql

import (
	"github.com/graphql-go/graphql"
)

func NewUrlSchema(query, mutation *graphql.Object) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query,
			Mutation: mutation,
		},
	)
}
