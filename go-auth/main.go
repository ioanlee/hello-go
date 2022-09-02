package main

import (
	"net/http"

	"github.com/ioanlee/hello-go/go-auth/users"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w, r)
	case "/sign-up-form":
		getSignUpPage(w, r)
	}
}

func getSignInPage(w http.ResponseWriter, r *http.Request) {}
func getSignUpPage(w http.ResponseWriter, r *http.Request) {}

func signInUser(w http.ResponseWriter, r *http.Request) {}
func signUpUser(w http.ResponseWriter, r *http.Request) {}

func getCredentials(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		email,
		password,
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("", nil)
}
