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

	oAuthSessionString = "ThisIsARandomKey"
)

// Environment Variable Settings
var (
	// Assembla OAuth Configuration
	assemblaClientID     = os.Getenv(assemblaClientIDKey)
	assemblaClientSecret = os.Getenv(assemblaClientSecretKey)

	// GitHub OAuth Configuration
	gitHubClientID     = os.Getenv(gitHubClientIDKey)
	gitHubClientSecret = os.Getenv(gitHubClientSecretKey)

	// General
	environment = os.Getenv("GO_ENV")
)

func main() {
	// Global Environment Variables
	if gitHubClientID == "" || gitHubClientSecret == "" {
		panic(fmt.Sprintf("%v or %s cannot be empty", gitHubClientIDKey, gitHubClientSecretKey))
	}

	if assemblaClientID == "" || assemblaClientSecret == "" {
		panic(fmt.Sprintf("%v or %s cannot be empty", assemblaClientIDKey, assemblaClientSecretKey))
	}

	if len(environment) == 0 {
		environment = "production"
	}

	// Database
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	if len(connectionString) == 0 {
		panic(fmt.Sprintf("%v cannot be empty", connectionString))
	}

	InitializeDatabase(connectionString)

	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// Routes
	// Login
	r.Handle("/login/{provider}", Login).Methods("GET")
	r.Handle("/login/{provider}/callback", LoginCallback).Methods("GET")

	//Users
	r.Handle("/users", UsersController).Methods("GET")
	r.Handle("/users/{userID}", UserController).Methods("GET")

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
