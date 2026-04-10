package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fantarqse/cacheserver/internal/core/model"
)

const key string = "url"

// TODO:
// 1) Read about JSON handling in the standard Go net/http library.
// 2) Read how to provide data to a server:
//   - base64
//   - Content as a binary file upload
//   - Raw bytes in request body

func (h *HTTPServer) Put(w http.ResponseWriter, r *http.Request) {
	log.Println("hit Put")

	url := r.URL.Query().Get(key)

	bytes, err := json.Marshal(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.service.Put(r.Context(), url, model.Page{Data: bytes}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
