package main

import (
	"html/template"
	"net/http"
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

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.verifyUser(newUser)
	fileName := "sign-in.html"
	t, _ := template.ParseFiles(fileName)
	if !ok {
		t.ExecuteTemplate(w, fileName, "user sign-in failure")
		return
	}
	t.ExecuteTemplate(w, fileName, "user sign-in success")
	return
}
func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.createUser(newUser)
	fileName := "sign-up.html"
	t, _ := template.ParseFiles(fileName)
	if err != nil {
		t.ExecuteTemplate(w, fileName, "new user sign-up failure")
		return
	}
	t.ExecuteTemplate(w, fileName, "new user sign-up success")
	return
}

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
