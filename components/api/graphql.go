package api

import (
	"net/http"

	"github.com/geowa4/base-go/components/api/foos"
	"github.com/geowa4/base-go/components/api/me"
	"github.com/geowa4/base-go/components/auth"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func newQueryType(db *sqlx.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me":   me.NewMeField(),
				"foos": foos.NewFooField(db),
			},
		},
	)
}

// NewGraphQLHandler creates a new HTTP handler for GraphQL.
func NewGraphQLHandler(db *sqlx.DB) http.Handler {
	queryType := newQueryType(db)
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
		ctx := auth.NewContextWithCurrentUser(r)
		graphQLHandler.ContextHandler(ctx, w, r)
	})
}
