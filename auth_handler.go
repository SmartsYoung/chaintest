package chaintest

import (
	"fmt"
	"github.com/SmartsYoung/chaintest/auth"
	"github.com/SmartsYoung/chaintest/oauth"
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
	"net/http"
)

const (
	name = "auth"
	Auth = "auth"
)

func init() {
	handler.HandlerFuncMap[Auth] = newAuthHandler
}

type AuthHandler struct {
}

// newAuthHandler fault handle gives the object of AuthHandler
func newAuthHandler() handler.Handler {
	return &AuthHandler{}
}

// Name function returns fault-inject string
func (ah *AuthHandler) Name() string {
	return "auth"
}

func (ah *AuthHandler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	if i.RouteTags.KV != nil {
		chain.Next(i, cb)
		return
	}

	fn := func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("Not found"))
			return
		}
		tokenStr := req.Header.Get("X-Auth-Token")
		resource := req.URL.String()
		token := &auth.Token{}
		token.Parse(tokenStr)

		auth ,err := auth.GetAuther(oauth.AuthName)
		if err == nil {
			fmt.Println("not found AuthName")
		}
		if err := auth.Authenticate(resource, token); err != nil {
			rw.Header().Set("X-Auth-Url", req.Host+"/v1/oauth/authorize")
			rw.WriteHeader(http.StatusForbidden)
			rw.Write([]byte("Forbidden"))
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("User is xxx"))
	}


	//call next chain
	chain.Next(i, cb)
}
