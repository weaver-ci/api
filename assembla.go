package main

import (
	"github.com/weaver-ci/models"
	"golang.org/x/oauth2"
)

// Endpoint is Assembla's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.assembla.com/authorization",
	TokenURL: "https://api.assembla.com/token",
}

// AssemblaOAuth implementation
type AssemblaOAuth struct {
	Configuration OAuthProperties
}

// NewAssemblaOAuth creates a AssemblaOAuth struct
func NewAssemblaOAuth(clientID, clientSecret, stateKey string) *AssemblaOAuth {
	a := new(AssemblaOAuth)

	a.Configuration.Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     Endpoint,
	}

	a.Configuration.StateKey = stateKey

	return a
}

// Login AssemblaOAuth
func (assembla AssemblaOAuth) Login() string {
	return assembla.Configuration.Config.AuthCodeURL(assembla.Configuration.StateKey, oauth2.AccessTypeOnline)
}

// Callback AssemblaOAuth
func (assembla AssemblaOAuth) Callback(state, code string) *oauth2.Token {
	return OAuthLoginCallback(assembla.Configuration.Config, assembla.Configuration.StateKey, state, code)
}

// GetUserInformation AssemblaOAuth
func (assembla AssemblaOAuth) GetUserInformation(token *oauth2.Token) *models.User {
	// Return a fake user till we access the Assembla API
	u := new(models.User)
	u.EmailAddress = "fake@assembla.com"

	return u
}
