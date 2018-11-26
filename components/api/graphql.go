package api

import (
	"context"
	"errors"
	"expvar"
	"net/http"

	"github.com/geowa4/base-go/components/api/foos"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var (
	countMeLoaded    = expvar.NewInt("NumLoadedUsers")
	stringLastLoaded = expvar.NewString("LastLoadedUser")
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
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
			user, ok := p.Context.Value(graphQLKey("currentUser")).(user)
			if !ok {
				return nil, errors.New("invalid type for 'me'")
			}
			countMeLoaded.Add(1)
			stringLastLoaded.Set(user.Name)
			return user, nil
		},
	}
}

// NewGraphQLHandler creates a new HTTP handler for GraphQL.
func NewGraphQLHandler(db *sqlx.DB) http.Handler {
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"me":   makeMeField(),
				"foos": foos.NewFooField(db),
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
		user := user{
			ID:   1,
			Name: "cool user",
		}
		ctx := context.WithValue(r.Context(), graphQLKey("currentUser"), user)
		graphQLHandler.ContextHandler(ctx, w, r)
	})
}
