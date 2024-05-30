package forumData

import (
	"database/sql"
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
		if _, err = db.Exec("INSERT INTO users (username, email, password, profile_picture, role) VALUES (?, ?, ?, ?, ?)", forumData[0], forumData[1], forumData[2], "Logo.png", "user"); err != nil {
			return err
		}
	case "postData":
		if _, err = db.Exec("INSERT INTO posts (uid, content, like, dislike) VALUES (?, ?, ?, ?)", forumData[0], forumData[1], forumData[2], forumData[4]); err != nil {
			return err
		}
	}
	return nil
}

func GetData(dataType string) ([]string, error) {
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
