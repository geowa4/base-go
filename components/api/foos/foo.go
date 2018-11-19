package foos

import (
	"github.com/jmoiron/sqlx"

	"github.com/graphql-go/graphql"
)

type foo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bars []*bar `json:"bars"`
}

type bar struct {
	ID    int  `json:"id"`
	Foo   *foo `json:"foo"`
	Value int  `json:"value"`
}

func makeFooType() *graphql.Object {
	var (
		fooType *graphql.Object
		barType *graphql.Object
	)
	barType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "bar",
			Fields: (graphql.FieldsThunk)(func() graphql.Fields {
				return graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.Int,
					},
					"value": &graphql.Field{
						Type: graphql.Int,
					},
					"foo": &graphql.Field{ // testing out cyclic types
						Type: fooType,
					},
				}
			}),
		},
	)
	fooType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "foo",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"bars": &graphql.Field{
					Type: graphql.NewList(barType),
				},
			},
		},
	)
	return fooType
}

//NewFooField makes a new GraphQL field for the foo type.
func NewFooField(db *sqlx.DB) *graphql.Field {
	fooType := makeFooType()
	resolver := &fooResolver{
		accessor: &fooDataAccessor{
			db: db,
		},
	}
	return &graphql.Field{
		Name: "foos",
		Type: graphql.NewList(fooType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.resolve,
	}
}
