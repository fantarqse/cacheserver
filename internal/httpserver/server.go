package httpserver

import (
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

func (s *HTTPServer) Serve() error {
	log.Println("serving...")

	s.mux.HandleFunc("POST /put", s.Put)
	s.mux.HandleFunc("GET /get", s.Get)
	s.mux.HandleFunc("DELETE /delete", s.Delete)
	s.mux.HandleFunc("GET /top", s.Top)

	return http.ListenAndServe(":8080", s.mux)
}
