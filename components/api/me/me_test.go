package me

import (
	"context"
	"testing"

	"github.com/geowa4/base-go/components/auth"

	"github.com/graphql-go/graphql"
)

func TestReturnsErrorIfNoUserSetOnContext(t *testing.T) {
	meResolver := NewMeField().Resolve
	user, err := meResolver(graphql.ResolveParams{
		Context: context.Background(),
	})
	if err == nil {
		t.Fail()
	}
	if user != auth.AnonymousUser {
		t.Fail()
	}
}

func TestPullsAnonymousUserFromTheContext(t *testing.T) {
	meResolver := NewMeField().Resolve
	testUser := auth.AnonymousUser
	userAsInterface, err := meResolver(graphql.ResolveParams{
		Context: context.WithValue(
			context.Background(),
			auth.CurrentUserKey,
			testUser,
		),
	})
	if err != nil {
		t.Error(err)
	}
	user, ok := userAsInterface.(auth.User)
	if !ok {
		t.Errorf("%v is not a User", userAsInterface)
	}
	if user.ID != testUser.ID {
		t.Errorf("%d is not %d", user.ID, testUser.ID)
	}
}

func TestPullsUserFromTheContext(t *testing.T) {
	meResolver := NewMeField().Resolve
	testUser := auth.User{ID: 1, Name: "Not Anonymous"}
	userAsInterface, err := meResolver(graphql.ResolveParams{
		Context: context.WithValue(
			context.Background(),
			auth.CurrentUserKey,
			testUser,
		),
	})
	if err != nil {
		t.Error(err)
	}
	user, ok := userAsInterface.(auth.User)
	if !ok {
		t.Errorf("%v is not a User", userAsInterface)
	}
	if user.ID != testUser.ID {
		t.Errorf("%d is not %d", user.ID, testUser.ID)
	}
}
