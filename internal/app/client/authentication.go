package client

import (
	"bytes"
	"encoding/json"
	cpaceLib "filippo.io/cpace"
	"fmt"
	"github.com/armr-dev/cpace-go/internal/app/cpace"
	"github.com/armr-dev/cpace-go/internal/utils"
	"net/http"
)

type ClientRegistration struct {
	MsgA     []byte
	UserName string
}

func authentication(username, password string) (string, error) {
	c := cpaceLib.NewContextInfo(username, string(cpace.ServerId), nil)

	msgA, state, err := cpaceLib.Start(password, c)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	clientReg := ClientRegistration{msgA, username}

	postBody, _ := json.Marshal(clientReg)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/authentication-init", "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var msgB []byte
	utils.Use(msgB, state)

	err = json.NewDecoder(resp.Body).Decode(&msgB)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	keyA, err := state.Finish(msgB)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	postBody, _ = json.Marshal(keyA)
	responseBody = bytes.NewBuffer(postBody)
	resp2, err := http.Post("http://localhost:8090/authentication-finalize", "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var respBody string

	err = json.NewDecoder(resp2.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return respBody, nil
}
