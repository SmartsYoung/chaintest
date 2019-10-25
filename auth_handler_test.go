package chaintest

import (
	"fmt"
	"github.com/SmartsYoung/chaintest/auth"
	"github.com/SmartsYoung/chaintest/oauth"
	"github.com/go-chassis/go-chassis/core/handler"
	"log"
	"net/http"
	"testing"
)

func TestAuthorize_Handle(t *testing.T) {

	var rw http.ResponseWriter
	var req *http.Request

	req, _ = http.NewRequest("POST", "http://localhost:8090/v1/oauth/authorize?username=SmartsYoung&password=hitsz_2019", nil)

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
	log.Println(token)
	/*rw.Header().Set("X-Auth-Token", token.Format())
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Success"))*/
}

func TestAuthHandler_Handle(t *testing.T) {

	authHandler := newAuthHandler()
	c := handler.Chain{}
	c.AddHandler(authHandler)

}
