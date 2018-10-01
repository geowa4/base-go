package api

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/zerolog/log"
)

type graphQLKey string

func makeMeField() *graphql.Field {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	return &graphql.Field{
		Type: userType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return p.Context.Value(graphQLKey("currentUser")), nil
		},
	}
}

// NewGraphQLHandler creates a new HTTP handler for GraphQL.
func NewGraphQLHandler() http.Handler {
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me": makeMeField(),
			},
		},
	)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("GraphQL initiailization failed.")
	}
	graphQLHandler := handler.New(&handler.Config{
		Schema: &schema,
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{1, "cool user"}
		ctx := context.WithValue(r.Context(), graphQLKey("currentUser"), user)
		graphQLHandler.ContextHandler(ctx, w, r)
	})
}
