package golibgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib/web/constant"
	"net/http"
)

type nextHandler struct{}

// Pull Gin's context from the request context and call the next item
// in the chain.
func (h *nextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	state := r.Context().Value(h).(*middlewareCtx)
	defer func(r *http.Request) { state.ctx.Request = r }(state.ctx.Request)
	state.childCalled = true
	state.ctx.Request = r
	state.ctx.Writer = &wrappedResponseWriter{state.ctx.Writer, w}
	state.ctx.Set(constant.ContextReqAttribute, r.Context().Value(constant.ContextReqAttribute))
	state.ctx.Next()
}

type middlewareCtx struct {
	ctx         *gin.Context
	childCalled bool
}

func New() (http.Handler, func(h http.Handler) gin.HandlerFunc) {
	next := new(nextHandler)
	return next, func(h http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			state := &middlewareCtx{ctx: c}
			ctx := context.WithValue(c.Request.Context(), next, state)
			h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
			if !state.childCalled {
				c.Abort()
			}
		}
	}
}

// Wrap takes the common HTTP middleware function signature, calls it to generate
// a handler, and wraps it into a Gin middleware handler.
//
// This is just a convenience wrapper around New.
func Wrap(f func(h http.Handler) http.Handler) gin.HandlerFunc {
	next, adapter := New()
	return adapter(f(next))
}

// WrapAll allow to wrap multiple http.Handler,
// returns a slice of gin.HandlerFunc
func WrapAll(hh []func(h http.Handler) http.Handler) []gin.HandlerFunc {
	functions := make([]gin.HandlerFunc, 0)
	for _, h := range hh {
		functions = append(functions, Wrap(h))
	}
	return functions
}
