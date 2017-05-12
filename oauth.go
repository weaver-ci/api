package main

import (
	"fmt"

	"golang.org/x/oauth2"

	"github.com/weaver-ci/models"
)

// OAuth2 interface
type OAuth2 interface {
	Login() string
	Callback(state, code string) *oauth2.Token
	GetUserInformation(token *oauth2.Token) *models.User
}

// OAuthProperties common configuration
type OAuthProperties struct {
	Config   oauth2.Config
	StateKey string
}

// LoginCallback generic login call back validator
func OAuthLoginCallback(configuration oauth2.Config, expectedState, state, code string) *oauth2.Token {
	if state != expectedState {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", expectedState, state)
		return nil
	}

	token, err := configuration.Exchange(oauth2.NoContext, code)

	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		return nil
	}

	return token
}
