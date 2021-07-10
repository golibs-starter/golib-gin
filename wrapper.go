package golibgin

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib/web/constants"
	"net/http"
)

// A wrapper that turns a http.ResponseWriter into a gin.ResponseWriter, given an existing gin.ResponseWriter
// Needed if the middleware you are using modifies the writer it passes downstream
// Wrap more methods: https://golang.org/pkg/net/http/#ResponseWriter
type wrappedResponseWriter struct {
	gin.ResponseWriter
	writer http.ResponseWriter
}

func (w *wrappedResponseWriter) Writer() http.ResponseWriter {
	return w.writer
}

func (w *wrappedResponseWriter) WriteString(s string) (n int, err error) {
	return w.writer.Write([]byte(s))
}

func (w *wrappedResponseWriter) WriteHeader(code int) {
	w.writer.WriteHeader(code)
}

// An http.Handler that passes on calls to downstream middlewares
type nextRequestHandler struct {
	c *gin.Context
}

// Run the next request in the middleware chain and return
func (h *nextRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.c.Set(constants.ContextReqAttribute, r.Context().Value(constants.ContextReqAttribute))
	h.c.Request = h.c.Request.WithContext(r.Context())
	h.c.Writer = &wrappedResponseWriter{h.c.Writer, w}
	h.c.Next()
}

// Wrap something that accepts an http.Handler, returns an gin.HandlerFunc
func Wrap(hh func(h http.Handler) http.Handler) gin.HandlerFunc {
	// Steps:
	// - create an http handler to pass `hh`
	// - call `hh` with the http handler, which returns a function
	// - call the ServeHTTP method of the resulting function to run the rest of the middleware chain

	return func(c *gin.Context) {
		hh(&nextRequestHandler{c}).ServeHTTP(c.Writer, c.Request)
	}
}

// WrapAll allow to wrap multiple http.Handler, returns a slice of gin.HandlerFunc
func WrapAll(hh []func(h http.Handler) http.Handler) []gin.HandlerFunc {
	functions := make([]gin.HandlerFunc, 0)
	for _, h := range hh {
		functions = append(functions, Wrap(h))
	}
	return functions
}
