package forumData

import "time"

type User struct {
	ID             int
	Username       string
	Email          string
	Password       string
	ProfilePicture string
	Role           string
	Post_nb        int
	Like_nb        int
	Dislike_nb     int
}

type Post struct {
	ID      int
	UID     int
	Text    string
	Image   string
	Like    int
	Dislike int
}

type Comment struct {
	ID      int
	UID     int
	PID     int
	Text    string
	Image   string
	Like    int
	Dislike int
}

type UserSession struct {
	Username   string
	ExpireTime time.Time
}
