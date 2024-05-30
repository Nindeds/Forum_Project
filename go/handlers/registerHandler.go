package handlers

import (
	"Forum_Project/go/forumData"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

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

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}
	IsDataUsed := isDataUsed(username, email)

	if isEmpty == false && IsDataUsed == false {
		err = forumData.InsertData("users", username, email, string(hashedPass))
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

func isDataUsed(data ...string) bool {
	emails, err := forumData.GetData("email")
	usernames, err := forumData.GetData("username")

	MailAlreadyUsed := false
	UsernameAlreadyUsed := false
	IsDataUsed := false

	if err != nil {
		panic(err)
	}

	for _, email := range emails {
		if email == data[0] {
			MailAlreadyUsed = true
		}
	}

	for _, username := range usernames {
		if username == data[1] {
			UsernameAlreadyUsed = true
		}
	}

	if UsernameAlreadyUsed == true || MailAlreadyUsed == true {
		IsDataUsed = true
		return IsDataUsed
	}

	return IsDataUsed
}

func isDataEmpty(data ...string) bool {
	for _, dataString := range data {
		if dataString == "" {
			return true
		}
	}
	return false
}
