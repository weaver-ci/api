package main

import (
	"database/sql"

	"github.com/weaver-ci/repository"
)

var (
	// UserRepository to manage Users
	UserRepository repository.UserRepository
)

// InitializeDatabase connects to the database and initializes all repositories
func InitializeDatabase(connectionString string) {
	db := repository.OpenDatabaseConnection(connectionString)

	initializeRepositories(db)
}

func initializeRepositories(db *sql.DB) {
	UserRepository = repository.NewPgUserRepository(db)
}
