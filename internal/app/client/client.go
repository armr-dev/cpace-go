package client

import "github.com/armr-dev/cpace-go/internal/app/cpace"

func InitClient() {
	registrationReq(cpace.DefaultUserName, cpace.DefaultPassword)
	authentication(cpace.DefaultUserName, cpace.DefaultPassword)
}
