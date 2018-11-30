package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

//ContextKey is used as keys for Contexts
type ContextKey string

//User represents the user making the request
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//CurrentUserKey is the Context key used for the current user.
const CurrentUserKey = ContextKey("currentUser")

//AnonymousUserID is the ID only an unauthenticated user can have.
const AnonymousUserID = 0

//AnonymousUser represents an unauthenticated user
var AnonymousUser = User{
	ID:   AnonymousUserID,
	Name: "Anonymous",
}

var errBadCredentials = fmt.Errorf("bad credentials used for authentication")

//TODO: this is fine for this template but not for any future use of it.
func newUserFromRequest(r *http.Request) (User, error) {
	authorization, ok := r.Header["Authorization"]
	if !ok {
		return AnonymousUser, errors.New("no credentials provided")
	} else if len(authorization) == 0 || authorization[0] == "" {
		return AnonymousUser, errBadCredentials
	}
	return User{
		ID:   1,
		Name: fmt.Sprintf("%v", authorization),
	}, nil
}

//NewContextWithCurrentUser parses the request to create a new context containing the current user.
//The current user may be the anonymous user if there is nothing in the request to make an identification.
func NewContextWithCurrentUser(r *http.Request) context.Context {
	reqCtx := r.Context()
	user, err := newUserFromRequest(r)
	if err == errBadCredentials {
		log.Info().Err(err).Msgf("%v", r.Header)
	}
	return context.WithValue(reqCtx, CurrentUserKey, user)
}
