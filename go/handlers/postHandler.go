package handlers

import (
	"Forum_Project/go/forumData"
	"html/template"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	UrlID := r.PathValue("id")
	PostData, err := forumData.GetSpecificPostData(UrlID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if PostData.Title == "" {
		http.Redirect(w, r, "/404", 303)
	}
	PostUserData, err := forumData.GetUserAndPostData(strconv.Itoa(PostData.UID), PostData)

	err = tmpl.Execute(w, PostUserData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
