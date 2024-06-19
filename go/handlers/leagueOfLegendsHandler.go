package handlers

import (
	"Forum_Project/go/forumData"
	"html/template"
	"net/http"
)

// LeagueOfLegendsHandler is the handler for the leagueoflegends category
func LeagueOfLegendsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/leagueoflegend.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postTitle := r.FormValue("title")
	postContent := r.FormValue("content")

	if postTitle != "" && postContent != "" {
		err = forumData.InsertData("postData", "27", postTitle, postContent)
		if err != nil {
			return
		}
	}

	AllPost, err := forumData.GetEveryPostData()
	if err != nil {
		return
	}

	err = tmpl.Execute(w, AllPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
