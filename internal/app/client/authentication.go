package client

import (
	"bytes"
	"encoding/json"
	cpaceLib "filippo.io/cpace"
	"fmt"
	"github.com/armr-dev/cpace-go/internal/app/cpace"
	"github.com/armr-dev/cpace-go/internal/utils"
	"log"
	"net/http"
)

func authentication() {
	c := cpaceLib.NewContextInfo(string(cpace.DefaultUserName), string(cpace.ServerId), nil)

	msgA, state, err := cpaceLib.Start(string(cpace.DefaultPassword), c)

	postBody, _ := json.Marshal(msgA)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/authentication-init", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var msgB []byte
	utils.Use(msgB, state)

	err = json.NewDecoder(resp.Body).Decode(&msgB)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	keyA, err := state.Finish(msgB)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	postBody, _ = json.Marshal(keyA)
	responseBody = bytes.NewBuffer(postBody)
	resp2, err := http.Post("http://localhost:8090/authentication-finalize", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var respBody string

	err = json.NewDecoder(resp2.Body).Decode(&respBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	fmt.Println(respBody)
}
