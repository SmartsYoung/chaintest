package chaintest

import (
	"fmt"
	"github.com/SmartsYoung/chaintest/auth"
	"github.com/SmartsYoung/chaintest/oauth"
	"log"
	"net/http"
	"testing"
)

func TestAuthHandler_Handle(t *testing.T) {

	var rw http.ResponseWriter
	var req *http.Request
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	log.Println(username, password)

	oauth, err := auth.GetAuther(oauth.AuthName)
	if err != nil {
		fmt.Println("not found AuthName")
		return
	}

	token, err := oauth.Authorize(username, password)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Unauthorized"))
		return
	}
	rw.Header().Set("X-Auth-Token", token.Format())
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Success"))

}
