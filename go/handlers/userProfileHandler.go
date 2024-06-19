package handlers

import (
	"Forum_Project/go/forumData"
	"fmt"
	"html/template"
	"net/http"
)

var UserData forumData.User

// UserProfileHandler is used to handle the userprofile page and show the profile of the inserted user
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	pathValue := r.PathValue("username")
	tmpl, err := template.ParseFiles("./src/html/userProfile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	UserData, err = forumData.GetAllUserData(pathValue)
	if UserData.Username == "" {
		http.Redirect(w, r, "/404", 303)
	}
	fmt.Println()

	err = tmpl.Execute(w, UserData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
