package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var (
	assemblaOAuth = NewAssemblaOAuth(assemblaClientID, assemblaClientSecret, oAuthSessionString)
	gitHubOAuth   = NewGithubOAuth(gitHubClientID, gitHubClientSecret, oAuthSessionString)
)

// Login utilizing oauth amongst different providers
var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var oauth OAuth2
	vars := mux.Vars(r)

	switch strings.ToLower(vars["provider"]) {
	case "assembla":
		oauth = assemblaOAuth
	case "github":
		oauth = gitHubOAuth
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url := oauth.Login()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
})

// LoginCallback utilizing oauth amongst different providers
var LoginCallback = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var oauth OAuth2
	vars := mux.Vars(r)

	state := r.FormValue("state")
	code := r.FormValue("code")

	switch strings.ToLower(vars["provider"]) {
	case "assembla":
		oauth = assemblaOAuth
		// handleAssemblaCallback(w, r)
	case "github":
		oauth = gitHubOAuth
		// handleGitHubCallback(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := oauth.Callback(state, code)

	if token == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u := oauth.GetUserInformation(token)

	if u == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeResponse(http.StatusOK, token, w, r)
})
