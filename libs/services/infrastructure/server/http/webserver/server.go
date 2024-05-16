package webserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	addr   string
}

func NewWebServer(addr string) *Server {
	return &Server{
		router: chi.NewRouter(),
		addr:   addr,
	}
}

// ConfigureDefaults sets up the default middleware for the server.
func (s *Server) ConfigureDefaults() {
	middlewares := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}
	s.RegisterMiddlewares(middlewares...)
}

// RegisterMiddlewares adds multiple middlewares to the server.
func (s *Server) RegisterMiddlewares(middlewares ...func(http.Handler) http.Handler) {
	for _, m := range middlewares {
		s.router.Use(m)
	}
}

// RegisterRoute adds a new route with an HTTP method, pattern, and handler function.
// If a group is specified, the route is added to that group.
func (s *Server) RegisterRoute(method, pattern string, handler http.HandlerFunc, group ...string) {
	if len(group) > 0 && group[0] != "" {
		r := s.router.Route(group[0], func(r chi.Router) {})
		r.MethodFunc(method, pattern, handler)
	} else {
		s.router.MethodFunc(method, pattern, handler)
	}
}

// RegisterRouteGroup registers a group of routes under a common prefix.
func (s *Server) RegisterRouteGroup(prefix string, routes func(r chi.Router)) {
	s.router.Route(prefix, routes)
}

// Start runs the web server on a specified address.
func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.router)
}
