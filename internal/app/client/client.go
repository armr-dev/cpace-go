package client

import "github.com/armr-dev/cpace-go/internal/app/cpace"

func InitClient() {
	Registration(cpace.DefaultUserName, cpace.DefaultPassword)
	Authentication(cpace.DefaultUserName, cpace.DefaultPassword)
}
