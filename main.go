package main

import (
	"Forum_Project/go/handlers"
	"fmt"
	"net/http"
)

type webserver struct {
	core      *http.ServeMux
	port      int
	assetsdir string
}

func main() {
	server := webserver{
		core:      http.NewServeMux(),
		port:      8080,
		assetsdir: "src",
	}
	server.Router()
	server.Launch()
}

func (s *webserver) Router() {
	s.core.HandleFunc("/register", handlers.RegisterHandler)
	s.core.HandleFunc("/login", handlers.LoginHandler)
	s.core.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./src/css"))))
	s.core.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./src/img"))))
	s.core.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./src/js"))))
	fmt.Printf("http://localhost:%d", s.port)
}

func (s *webserver) Launch() {
	// Lancement du serveur HTTP sur le port spécifié
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.core)
}
