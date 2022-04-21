package server

import (
	"github.com/armr-dev/cpace-go/internal/app/db"
	"net/http"
)

var userRecords = db.Record{}
var k = Keys{}

func InitServer() {
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/authentication-init", k.authenticationInit)
	http.HandleFunc("/authentication-finalize", k.authenticationFinalize)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
