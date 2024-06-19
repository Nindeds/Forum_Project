package handlers

import (
	"Forum_Project/go/forumData"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/updateProfile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prevUsername := r.FormValue("")
	newUsername := r.FormValue("username")
	newEmail := r.FormValue("email")
	newPassword := r.FormValue("password")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		fmt.Println(err)
	}

	forumData.UpdateUserData(prevUsername, newUsername, newEmail, string(hashedPass), "Logo.png", false)

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
