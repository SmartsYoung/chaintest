package oauth

import (
	"context"
	"encoding/json"
	"github.com/SmartsYoung/chaintest/auth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"strings"
)

const (
	SecretKey = "welcome!"
)

var AuthName = "oauth2"

type OAuth struct {
}

func init() {
	auth.RegisterAuther(AuthName, &OAuth{})
}

var(
	key []byte = []byte("Hello World!")
)


func (o *OAuth) Authorize(username, password string) (*auth.Token, error) {
	conf := oauth2.Config{
		ClientID: "0a41d2e6b263a74a55c8eebd3ba17bdc67b13445ec3456edec409749a92ad9d7",
		ClientSecret:"f3077ede05ab96ac5187fe627e7b9fb3080d755d55b94af628a3a568bc5114d1",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://gitee.com/oauth/authorize",
			TokenURL: "https://gitee.com/oauth/token",
			AuthStyle: 0,
		},
	}

	token, err := conf.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	b, _ := json.Marshal(token)

	log.Println(string(b))

	return nil, nil
}

func (o *OAuth) Authenticate(source string, token *auth.Token) error {
	if strings.Compare(source, token.AccessToken) == -1 {
		return nil
	}
	return nil
}


func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
