package utils

import (
	"errors"
	"github.com/armr-dev/cpace-go/internal/app/user"
)

// Used to ignore declared and not used annoying message
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func FindUser(users []user.User, userName string) (user.User, error) {
	for i := range users {
		if string(users[i].UserName) == userName {
			return users[i], nil
		}
	}
	return user.User{}, errors.New("user not found")
}
