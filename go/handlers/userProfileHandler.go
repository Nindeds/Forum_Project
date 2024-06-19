package handlers

import (
	"Forum_Project/go/forumData"
	"html/template"
	"net/http"
)

var UserData forumData.User

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	pathValue := r.PathValue("username")
	tmpl, err := template.ParseFiles("./src/html/userProfile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	UserData, err = forumData.GetAllUserData(pathValue)

	err = tmpl.Execute(w, UserData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
