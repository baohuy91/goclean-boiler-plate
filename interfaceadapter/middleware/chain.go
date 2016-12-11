package middleware

import (
	"net/http"
)

type ChainFunc func(h http.Handler) http.Handler

// These are middleware to support http request
type Chain struct {
	chainFunctions []ChainFunc
}

// Factory function for new middleware
func NewChain(chainFunctions...ChainFunc) Chain {
	return Chain{
		chainFunctions:chainFunctions,
	}
}

func (c Chain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range c.chainFunctions {
		h = c.chainFunctions[len(c.chainFunctions)-1-i](h)
	}

	return h
}

func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then(nil)
	}
	return c.Then(fn)
}






