package server

import (
	"encoding/json"
	cpaceLib "filippo.io/cpace"
	"github.com/armr-dev/cpace-go/internal/app/client"
	"github.com/armr-dev/cpace-go/internal/app/cpace"
	"github.com/armr-dev/cpace-go/internal/utils"
	"net/http"
)

func (k *Keys) AuthenticationInit(w http.ResponseWriter, req *http.Request) {
	var clientReg client.ClientRegistration
	var err = json.NewDecoder(req.Body).Decode(&clientReg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := utils.FindUser(userRecords.Users, clientReg.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := cpaceLib.NewContextInfo(string(user.UserName), string(cpace.ServerId), nil)

	msgB, keyB, err := cpaceLib.Exchange(string(user.Password), c, clientReg.MsgA)

	err = json.NewEncoder(w).Encode(msgB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	k.KeyB = keyB
}

func (k *Keys) AuthenticationFinalize(w http.ResponseWriter, req *http.Request) {
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

	//fmt.Println(response)
}
