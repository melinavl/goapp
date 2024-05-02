package httpsrv

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/csrf"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	HFunc   http.Handler
	Queries []string
}

func (s *Server) myRoutes() []Route {
	return []Route{
		{
			Name:    "health",
			Method:  "GET",
			Pattern: "/goapp/health",
			HFunc:   s.handlerWrapper(s.handlerHealth),
		},
		{
			Name:    "websocket",
			Method:  "GET",
			Pattern: "/goapp/ws",
			HFunc:   s.handlerWrapper(s.handlerWebSocket),
		},
		{
			Name:    "home",
			Method:  "GET",
			Pattern: "/goapp",
			HFunc:   s.handlerWrapper(s.handlerHome),
		},
	}
}

func (s *Server) handlerWrapper(handlerFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	// Initialize CSRF protection with a secret key.This key should be kept in a secret or an env var.
	csrfMiddleware := csrf.Protect([]byte("nc98P987bcpncYhoadjoiydc9ajDlcn"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				s.error(w, http.StatusInternalServerError, fmt.Errorf("%v\n%v", r, string(debug.Stack())))
			}
		}()

		// Wrap handler with the CSRF middleware.
		csrfMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerFunc(w, r)
		})).ServeHTTP(w, r)
	})
}
