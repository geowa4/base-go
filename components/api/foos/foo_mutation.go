package foos

import (
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

//The cyclic types I chose for this example can make these mutations a little
//complicated so I'm making a simple type that just returns the ID.
var simpleCreateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Creation",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

//NewFooCreateField makes a new GraphQL field for creating the foo type.
func NewFooCreateField(db *sqlx.DB) *graphql.Field {
	return &graphql.Field{
		Name: "foos",
		Type: simpleCreateType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			resolver := newCreateResolver(db)
			return resolver.createFoo(params.Args["name"].(string))
		},
	}
}

//NewBarCreateField makes a new GraphQL field for creating the bar type.
func NewBarCreateField(db *sqlx.DB) *graphql.Field {
	return &graphql.Field{
		Name: "bars",
		Type: simpleCreateType,
		Args: graphql.FieldConfigArgument{
			"fooID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"value": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			resolver := newCreateResolver(db)
			return resolver.createBar(params.Args["fooID"].(int), params.Args["value"].(int))
		},
	}
}
