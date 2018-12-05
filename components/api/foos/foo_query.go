package foos

import (
	"github.com/jmoiron/sqlx"

	"github.com/graphql-go/graphql"
)

//NewFooQueryField makes a new GraphQL field for querying the foo type.
func NewFooQueryField(db *sqlx.DB) *graphql.Field {
	fooType := makeFooType()
	return &graphql.Field{
		Name: "foos",
		Type: graphql.NewList(fooType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			resolver := newLoadResolver(db)
			return resolver.query(params.Args, params.Info.FieldASTs)
		},
	}
}
