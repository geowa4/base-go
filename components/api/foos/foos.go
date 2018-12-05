package foos

import "github.com/graphql-go/graphql"

type foo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bars []*bar `json:"bars"`
}

type bar struct {
	ID    int64 `json:"id"`
	Foo   *foo  `json:"foo"`
	Value int64 `json:"value"`
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
