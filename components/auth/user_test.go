package auth

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReturnAnonymousUserIfNoAuthorizationHeader(t *testing.T) {
	mockRequest := httptest.NewRequest("GET", "/", nil)
	user, err := newUserFromRequest(mockRequest)
	if err == nil {
		t.Error(err)
	}
	if user.ID != AnonymousUserID {
		t.Errorf("%d is not %d", user.ID, AnonymousUserID)
	}
}

func TestReturnUserForAuthorization(t *testing.T) {
	mockRequest := httptest.NewRequest("GET", "/", nil)
	authorization := "token"
	mockRequest.Header.Add("Authorization", authorization)
	user, err := newUserFromRequest(mockRequest)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(user.Name, authorization) {
		t.Errorf("%s is not in %s", authorization, user.Name)
	}
}

func TestReturnBadCredentialsErrorForEmptyAuthorization(t *testing.T) {
	mockRequest := httptest.NewRequest("GET", "/", nil)
	authorization := ""
	mockRequest.Header.Add("Authorization", authorization)
	user, err := newUserFromRequest(mockRequest)
	if err != errBadCredentials {
		t.Error(err)
	}
	if user.ID != AnonymousUserID {
		t.Errorf("%d is not %d", user.ID, AnonymousUserID)
	}
}

func TestContextWithAnonymousUser(t *testing.T) {
	mockRequest := httptest.NewRequest("GET", "/", nil)
	ctx := NewContextWithCurrentUser(mockRequest)
	user, ok := ctx.Value(CurrentUserKey).(User)
	if !ok {
		t.Fail()
	}
	if user.ID != AnonymousUserID {
		t.Errorf("%d is not %d", user.ID, AnonymousUserID)
	}
}

func TestContextWithGoodUser(t *testing.T) {
	mockRequest := httptest.NewRequest("GET", "/", nil)
	authorization := "token"
	mockRequest.Header.Add("Authorization", authorization)
	ctx := NewContextWithCurrentUser(mockRequest)
	user, ok := ctx.Value(CurrentUserKey).(User)
	if !ok {
		t.Fail()
	}
	if user.ID == AnonymousUserID {
		t.Error("Got anonymous user.")
	}
}
