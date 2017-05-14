package main

import (
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// UsersController handle /users
var UsersController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := UserRepository.GetUsers()

		if err != nil {
			handleError(err, w, r)
			return
		}

		encodeResponse(http.StatusOK, users, w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
})

var UserController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := uuid.FromString(vars["userID"])

	if err != nil {
		encodeResponse(http.StatusBadRequest, nil, w, r)
		return
	}

	switch r.Method {
	case "GET":
		user, err := UserRepository.GetUser(userID)

		if err != nil {
			handleError(err, w, r)
			return
		}

		encodeResponse(http.StatusOK, user, w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
})
