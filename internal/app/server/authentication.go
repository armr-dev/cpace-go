package server

import (
	"encoding/json"
	cpaceLib "filippo.io/cpace"
	"fmt"
	"github.com/armr-dev/cpace-go/internal/app/cpace"
	"net/http"
)

func (k *Keys) authenticationInit(w http.ResponseWriter, req *http.Request) {
	var user = userRecords.Users[0]
	c := cpaceLib.NewContextInfo(string(user.UserName), string(cpace.ServerId), nil)
	var msgA []byte

	var err = json.NewDecoder(req.Body).Decode(&msgA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msgB, keyB, err := cpaceLib.Exchange(string(user.Password), c, msgA)

	err = json.NewEncoder(w).Encode(msgB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	k.KeyB = keyB
}

func (k *Keys) authenticationFinalize(w http.ResponseWriter, req *http.Request) {
	var keyA []byte

	var err = json.NewDecoder(req.Body).Decode(&keyA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	k.KeyA = keyA

	var response = "User not authenticated"

	if k.keysAreEqual() {
		response = "User authenticated"
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(response)
}
