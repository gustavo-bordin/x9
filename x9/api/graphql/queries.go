package graphql

import "github.com/graphql-go/graphql"

type Query struct {
	resolver Resolver
}

func NewQuery(resolver Resolver) Query {
	query := Query{
		resolver: resolver,
	}
	return query
}

func (q Query) GetUrlQueries() *graphql.Object {
	queryConfig := graphql.ObjectConfig{
		Name: "UrlQueries",
		Fields: graphql.Fields{
			"urls": &graphql.Field{
				Type:    graphql.NewList(urlType),
				Resolve: q.resolver.urls,
			},
			"url": &graphql.Field{
				Type: urlType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: q.resolver.urlById,
			},
		},
	}

	obj := graphql.NewObject(queryConfig)
	return obj
}

func (q Query) GetUrlMutations() *graphql.Object {
	queryConfig := graphql.ObjectConfig{
		Name: "UrlMutation",
		Fields: graphql.Fields{
			"addUrl": &graphql.Field{
				Type: urlType,
				Args: graphql.FieldConfigArgument{
					"url": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: q.resolver.addUrl,
			},
			"deleteUrl": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: q.resolver.deleteUrl,
			},
		},
	}

	obj := graphql.NewObject(queryConfig)
	return obj
}
