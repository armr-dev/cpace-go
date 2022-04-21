package server

import (
	"github.com/armr-dev/cpace-go/internal/app/db"
	"net/http"
)

var userRecords = db.Record{}

func InitServer() {
	http.HandleFunc("/registration-init", registrationInit)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
