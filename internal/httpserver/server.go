package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/fantarqse/cacheserver/internal/core/port"
)

type HTTPServer struct {
	mux     *http.ServeMux
	service port.CacheService
}

func New(service port.CacheService) *HTTPServer {
	return &HTTPServer{mux: http.NewServeMux(), service: service}
}

func (s *HTTPServer) Serve(_ context.Context, port int) error {
	log.Printf("serving at %d\n", port)

	s.mux.HandleFunc("POST /put", s.Put)
	s.mux.HandleFunc("GET /get", s.Get)
	s.mux.HandleFunc("DELETE /delete", s.Delete)
	s.mux.HandleFunc("GET /top", s.Top)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.mux)
}
