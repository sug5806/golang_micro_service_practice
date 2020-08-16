package oauth

import (
	"fmt"
	"golang_micro_service_practice/api/utils/errors"
)

const (
	getUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"fede": &User{
			Id:       1,
			Username: "fede",
		},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user found with given parameters"))
	}

	return user, nil
}
