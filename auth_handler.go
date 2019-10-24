package chaintest

import (
	//"fmt"
	//"github.com/SmartsYoung/chaintest/auth"
	//"github.com/SmartsYoung/chaintest/oauth"
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
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

	//call next chain
	chain.Next(i, cb)
}
