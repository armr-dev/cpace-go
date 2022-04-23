package client

import (
	"bytes"
	"encoding/json"
	"github.com/armr-dev/cpace-go/internal/app/user"
	"net/http"
)

func Registration(username, password string) (string, error) {
	var newUser = user.User{
		UserName: []byte(username),
		Password: []byte(password),
	}

	postBody, _ := json.Marshal(newUser)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/registration", "application/json", responseBody)
	if err != nil {
		return "", err
	}

	var respBody string

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return "", err
	}

	return respBody, nil
}
