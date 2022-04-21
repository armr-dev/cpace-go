package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/armr-dev/cpace-go/internal/app/cpace"
	"github.com/armr-dev/cpace-go/internal/app/user"
	"log"
	"net/http"
)

func registrationReq() {
	var newUser = user.User{
		cpace.DefaultUserName,
		cpace.DefaultPassword,
	}

	postBody, _ := json.Marshal(newUser)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/registration-init", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var respBody string

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	fmt.Println(respBody)
}
