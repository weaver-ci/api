package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"fmt"
)

const (
	assemblaClientIDKey     = "ASSEMBLA_CLIENT_ID"
	assemblaClientSecretKey = "ASSEMBLA_CLIENT_SECRET"
	gitHubClientIDKey       = "GITHUB_CLIENT_ID"
	gitHubClientSecretKey   = "GITHUB_CLIENT_SECRET"
)

var (
	// Assembla OAuth Configuration
	assemblaClientID     = os.Getenv(assemblaClientIDKey)
	assemblaClientSecret = os.Getenv(assemblaClientSecretKey)

	// GitHub OAuth Configuration
	gitHubClientID     = os.Getenv(gitHubClientIDKey)
	gitHubClientSecret = os.Getenv(gitHubClientSecretKey)

	environment = os.Getenv("GO_ENV")

	oAuthSessionString = "ThisIsARandomKey"
)

func main() {
	// Validate Properties
	if gitHubClientID == "" || gitHubClientSecret == "" {
		panic(fmt.Sprintf("%v or %s cannot be empty", gitHubClientIDKey, gitHubClientSecretKey))
	}

	if assemblaClientID == "" || assemblaClientSecret == "" {
		panic(fmt.Sprintf("%v or %s cannot be empty", assemblaClientIDKey, assemblaClientSecretKey))
	}

	if len(environment) == 0 {
		environment = "production"
	}

	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// Routes
	r.Handle("/login/{provider}", Login).Methods("GET")
	r.Handle("/login/{provider}/callback", LoginCallback).Methods("GET")

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
