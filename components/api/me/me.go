package me

import (
	"errors"
	"expvar"

	"github.com/geowa4/base-go/components/auth"
	"github.com/graphql-go/graphql"
)

var (
	countMeLoaded    = expvar.NewInt("NumLoadedUsers")
	stringLastLoaded = expvar.NewString("LastLoadedUser")
)

//NewMeField returns simple information about the current user.
func NewMeField() *graphql.Field {
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
			user, ok := p.Context.Value(auth.CurrentUserKey).(auth.User)
			if !ok {
				return auth.AnonymousUser, errors.New("invalid type for 'me'")
			}
			countMeLoaded.Add(1)
			stringLastLoaded.Set(user.Name)
			return user, nil
		},
	}
}
