package db

import (
	"github.com/armr-dev/cpace-go/internal/app/user"
)

type Record struct {
	Users []user.User `default:"[]"`
}
