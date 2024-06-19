package forumData

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

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
		_, err = db.Exec("INSERT INTO post (uid, content, like, dislike) VALUES (?, ?, ?, ?)", forumData[0], forumData[1], forumData[2], forumData[3])
		if err != nil {
			return err
		}
	}
	return nil
}

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

func GetAllUserData(username string) (User, error) {
	var userData User
	db, err := sql.Open("sqlite3", "db.db")

	defer db.Close()

	if err != nil {
		fmt.Println("pates")
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
