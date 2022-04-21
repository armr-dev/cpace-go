package server

import (
	"encoding/json"
	"fmt"
	"github.com/armr-dev/cpace-go/internal/app/user"
	"net/http"
)

func registration(w http.ResponseWriter, req *http.Request) {
	var newUser user.User

	var err = json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode("User registered!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRecords.Users = append(userRecords.Users, newUser)
	fmt.Println("User registered successfully!")
}
