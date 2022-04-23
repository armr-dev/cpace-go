package server

import (
	"github.com/armr-dev/cpace-go/internal/app/db"
	"net/http"
)

var userRecords = db.Record{}
var k = Keys{}

func InitServer() {
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/authentication-init", k.AuthenticationInit)
	http.HandleFunc("/authentication-finalize", k.AuthenticationFinalize)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
