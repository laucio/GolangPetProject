package Entity

import "errors"

type Project struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	EmptyRepo   bool   `json:"empty_repo"`
	ReadmeUrl   string `json:"readme_url"`
	StartDate   string `json:"created_at"`
	EndDate     string `json:"last_activity_at"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	User{Username: "laucio", Password: "lauciolaucio2"},
	User{Username: "user2", Password: "pass2"},
	User{Username: "user3", Password: "pass3"},
}

func registerNewUser(username, password string) (*User, error) {
	return nil, errors.New("placeholder error")
}

func isUsernameAvailable(username string) bool {
	return false
}
