package httpserver

import (
	"log"
	"net/http"
)

func (h *HTTPServer) Put(w http.ResponseWriter, r *http.Request) {
	log.Println("hit Put")
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *HTTPServer) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("hit Get")
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *HTTPServer) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("hit Delete")
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *HTTPServer) Top(w http.ResponseWriter, r *http.Request) {
	log.Println("hit Top")
	w.WriteHeader(http.StatusNotImplemented)
}
