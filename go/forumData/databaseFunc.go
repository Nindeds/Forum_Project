package forumData

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// CreateUserTable create a user table if it doesn't already
func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return err
	}

	db.Exec("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY ON CONFLICT FAIL AUTOINCREMENT NOT NULL ON CONFLICT FAIL, username TEXT    NOT NULL ON CONFLICT FAIL,\n\t\temail           TEXT    NOT NULL ON CONFLICT FAIL, password TEXT NOT NULL ON CONFLICT FAIL,profile_picture TEXT, role TEXT NOT NULL ON CONFLICT FAIL DEFAULT user, post_nb INTEGER NOT NULL DEFAULT (0), like_nb INTEGER NOT NULL DEFAULT (0), dislike_nb INTEGER NOT NULL DEFAULT (0))")
	return err
}

// InsertData use the dataType to determine in which table he inserts the data into
func InsertData(dataType string, forumData ...string) error {
	db, err := sql.Open("sqlite3", "db.db")

	defer db.Close()

	if err != nil {
		return err
	}
	switch dataType {
	case "userData":
		_, err = db.Exec("INSERT INTO users (username, email, password, profile_picture, role) VALUES (?, ?, ?, ?, ?)", forumData[0], forumData[1], forumData[2], "Logo.png", "user")
		if err != nil {
			return err
		}
	case "postData":
		_, err = db.Exec("INSERT INTO post (uid, title, content, like, dislike) VALUES (?, ?, ?, ?, ?)", forumData[0], forumData[1], forumData[2], 0, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSpecificUserData get a certain data from all the users of the users table
func GetSpecificUserData(dataType string) ([]string, error) {
	var datas []string
	var data string
	db, err := sql.Open("sqlite3", "db.db")

	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, _ := db.Query("SELECT " + dataType + " FROM users")
	for rows.Next() {
		err = rows.Scan(&data)
		if err != nil {
			return datas, err
		}
		datas = append(datas, data)
	}

	return datas, err
}

// UpdateUserData is used to update the data in the user table when the user change his data or when the admin change a role
func UpdateUserData(prevUsername string, username string, email string, hashedpassword string, profilepicture string, isAdmin bool) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return
	}
	switch isAdmin {
	case false:
		_, err = db.Exec("UPDATE users SET username=?, email=?, password=?, profile_picture=? WHERE username=?", username, email, hashedpassword, profilepicture, prevUsername)
	case true:
		_, err = db.Exec("UPDATE users SET role=? WHERE username=?", username, email, hashedpassword, profilepicture, prevUsername)

	}
	return
}

// GetAllUserData get all the data from a specific user
func GetAllUserData(username string) (User, error) {
	var userData User
	db, err := sql.Open("sqlite3", "db.db")

	defer db.Close()

	if err != nil {
		return userData, err
	}

	rows, _ := db.Query("SELECT * FROM users WHERE username='" + username + "'")
	for rows.Next() {
		err = rows.Scan(&userData.ID, &userData.Username, &userData.Email, &userData.Password, &userData.ProfilePicture, &userData.Role, &userData.Like_nb, &userData.Dislike_nb, &userData.Post_nb)
		if err != nil {
			fmt.Println(err)
			return userData, err
		}
	}

	return userData, err
}

// GetEveryPostData get all the data from every post and return them as a list
func GetEveryPostData() ([]Post, error) {
	var postData Post
	var allPostData []Post
	db, err := sql.Open("sqlite3", "db.db")

	rows, _ := db.Query("SELECT * FROM post")
	for rows.Next() {
		err = rows.Scan(&postData.ID, &postData.UID, &postData.Title, &postData.Text, &postData.Like, &postData.Dislike)
		if err != nil {
			return allPostData, err
		}
		allPostData = append(allPostData, postData)
	}
	return allPostData, err
}

// GetSpecificPostData get the data of a post from the id
func GetSpecificPostData(id string) (Post, error) {
	var postData Post
	db, err := sql.Open("sqlite3", "db.db")

	rows, _ := db.Query("SELECT * FROM post WHERE ID=" + id)
	for rows.Next() {
		err = rows.Scan(&postData.ID, &postData.UID, &postData.Title, &postData.Text, &postData.Like, &postData.Dislike)
		if err != nil {
			return postData, err
		}
	}
	return postData, err
}

// GetUserAndPostData is used to get both the data of a post with the username of the person who created the post
func GetUserAndPostData(uid string, postData Post) (ShowPost, error) {

	var username string
	var showPost ShowPost

	db, err := sql.Open("sqlite3", "db.db")
	rows, err := db.Query("SELECT username FROM users WHERE ID=" + uid)
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return showPost, err
		}
	}

	showPost.Like = postData.Like
	showPost.Text = postData.Text
	showPost.Title = postData.Title
	showPost.Dislike = postData.Dislike
	showPost.Username = username
	return showPost, err
}
