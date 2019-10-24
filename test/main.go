package main

import (
	"fmt"
	"github.com/SmartsYoung/chaintest/auth"
	"github.com/SmartsYoung/chaintest/oauth"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	http.HandleFunc("/v1/oauth/authorize", authorize)
	http.HandleFunc("/v1/user", getUser)
	http.ListenAndServe(":8090", nil)

}

func getUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Not found"))
		return
	}
	tokenStr := r.Header.Get("X-Auth-Token")
	resource := r.URL.String()
	token := &auth.Token{}
	token.Parse(tokenStr)

	auth ,err := auth.GetAuther(oauth.AuthName)
	if err == nil {
		fmt.Println("not found AuthName")
	}
	if err := auth.Authenticate(resource, token); err != nil {
		rw.Header().Set("X-Auth-Url", r.Host+"/v1/oauth/authorize")
		rw.WriteHeader(http.StatusForbidden)
		rw.Write([]byte("Forbidden"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User is xxx"))

}



func authorize(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Not found"))
		return
	}
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	log.Println(username, password)

	oauth ,err := auth.GetAuther(oauth.AuthName)
	if err != nil {
		fmt.Println("not found AuthName")
		return
	}

	token, err := oauth.Authorize(username, password )
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Unauthorized"))
		return
	}
	rw.Header().Set("X-Auth-Token", token.Format())
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Success"))
}
