package handlers

import (
	"Forum_Project/go/forumData"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

var sessionMap map[string]forumData.UserSession

type cError struct {
	WrongCredentials bool
}

// LoginHandler is the handler for the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookieSession, err := r.Cookie("userSession")
	if cookieSession != nil {
		http.Redirect(w, r, "/", 303)
	}

	tmpl, err := template.ParseFiles("./src/html/Login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	UserError := cError{WrongCredentials: false}

	username_email := r.FormValue("User")
	password := r.FormValue("UserPassword")
	username, isCorrect, err := isLogInfoCorrect(username_email, password)

	// This if is used to give the cookie to the user and check whether or not the user inserted the good credentials
	if isCorrect == false && password != "" && username != "" {
		UserError = cError{WrongCredentials: true}
	} else if isCorrect == true && cookieSession == nil {
		expiresAt := time.Now().Add(1 * time.Hour)

		var data = forumData.UserSession{
			Username:   username,
			ExpireTime: expiresAt,
		}

		uuid4, err := uuid.NewV4()
		if err != nil {

		}
		uuidString := fmt.Sprintf("%v", uuid4)

		sessionMap[uuidString] = data

		http.SetCookie(w, &http.Cookie{
			Name:  "userSession",
			Value: uuidString,

			Expires: expiresAt,
		})
	}

	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, UserError)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// isLogInfoCorrect is to check if the username/email is the right one and check the password of the said user
func isLogInfoCorrect(userInfo string, userPassword string) (string, bool, error) {
	var rows *sql.Rows
	var password, username string
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return "", false, err
	}
	hasAtChar := strings.Contains(userInfo, "@")

	if hasAtChar {
		rows, err = db.Query("SELECT password, username FROM users WHERE email='" + userInfo + "'")
	} else {
		rows, err = db.Query("SELECT password FROM users WHERE username='" + userInfo + "'")
	}

	for rows.Next() {
		if hasAtChar {
			err = rows.Scan(&password, &username)
		} else {
			err = rows.Scan(&password)
		}

		if err != nil {
			return "", false, err
		}
	}
	bytePassword := []byte(userPassword)
	byteHash := []byte(password)

	if bcrypt.CompareHashAndPassword(byteHash, bytePassword) == nil && hasAtChar == true {
		return username, true, err
	} else if bcrypt.CompareHashAndPassword(byteHash, bytePassword) == nil && hasAtChar == false {
		return userInfo, true, err
	}

	return "", false, err
}
