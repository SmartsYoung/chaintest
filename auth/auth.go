package auth

import (
	"fmt"
	"log"
	"time"
)

// HandlerFuncMap handler function map
var autherMap = map[string]Auther{}

type Auther interface {
	Authorize(username, password string) (*Token, error)
	Authenticate(source string, token *Token) error
}

func RegisterAuther(kind string, auther Auther) {
	_, ok := autherMap[kind]
	if ok {
		log.Printf("auther is already exit, name = %s", kind)
		return
	}
	autherMap[kind] = auther
}

func GetAuther(kind string) (Auther, error) {
	a, ok := autherMap[kind]
	if !ok {
		return nil, fmt.Errorf("auther is not found, name = %s", kind)
	}
	return a, nil
}

type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

func (T *Token) Parse(str string) {

}

func (T *Token) Format() string {
	s := T.AccessToken
	return s
}
