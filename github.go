package main

import (
	"fmt"

	"github.com/weaver-ci/models"

	githubclient "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

// GithubOAuth implementation
type GithubOAuth struct {
	Configuration OAuthProperties
}

// NewGithubOAuth creates a GitHubOAuth struct
func NewGithubOAuth(clientID, clientSecret, stateKey string) *GithubOAuth {
	g := new(GithubOAuth)

	g.Configuration.Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// select level of access you want https://developer.github.com/v3/oauth/#scopes
		Scopes:   []string{"user:email", "repo"},
		Endpoint: githuboauth.Endpoint,
	}

	g.Configuration.StateKey = stateKey

	return g
}

// Login GitHubOAuth
func (github GithubOAuth) Login() string {
	return github.Configuration.Config.AuthCodeURL(github.Configuration.StateKey, oauth2.AccessTypeOnline)
}

// Callback GithubOAuth
func (github GithubOAuth) Callback(state, code string) *oauth2.Token {
	return OAuthLoginCallback(github.Configuration.Config, github.Configuration.StateKey, state, code)
}

// GetUserInformation GithubOAuth
func (github GithubOAuth) GetUserInformation(token *oauth2.Token) *models.User {
	oauthClient := github.Configuration.Config.Client(oauth2.NoContext, token)
	client := githubclient.NewClient(oauthClient)

	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		return nil
	}

	u := new(models.User)
	u.EmailAddress = *user.Login
	return u
}
