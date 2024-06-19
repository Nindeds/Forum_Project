package handlers

import (
	"Forum_Project/go/forumData"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// RegisterHandler is used to handle the register page
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/Register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.FormValue("Username")
	email := r.FormValue("UserMail")
	password := r.FormValue("UserPassword")

	isEmpty := isDataEmpty(username, email, password)
	IsDataUsed := isDataUsed(email, username)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}

	if isEmpty == false && IsDataUsed == false {
		err = forumData.InsertData("userData", username, email, string(hashedPass))
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/login", 302)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// isDataUsed is used to verify if the inserted data is already used by another user
func isDataUsed(emailInput string, usernameInput string) bool {
	emails, err := forumData.GetSpecificUserData("email")
	usernames, err := forumData.GetSpecificUserData("username")

	MailAlreadyUsed := false
	UsernameAlreadyUsed := false
	IsDataUsed := false

	if err != nil {
		panic(err)
	}

	for _, email := range emails {
		if email == emailInput {
			MailAlreadyUsed = true
		}
	}

	for _, username := range usernames {
		if username == usernameInput {
			UsernameAlreadyUsed = true
		}
	}

	if UsernameAlreadyUsed == true || MailAlreadyUsed == true {
		IsDataUsed = true
		return IsDataUsed
	}

	return IsDataUsed
}

// isDataEmpty Check if the inserted data is empty
func isDataEmpty(username string, email string, password string) bool {
	if username == "" || email == "" || password == "" {
		return true
	}
	return false
}
