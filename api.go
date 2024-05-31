package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(Addr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: Addr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/account", customFuncToHandlerFunc(s.handleGetAccount))
	log.Println("Server running on port ", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}

type APIError struct {
	Error string
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func customFuncToHandlerFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	user := newAccount("Tilekbergen", "Mukhamet")
	if r.Method == "GET" {
		return writeJSON(w, http.StatusOK, user)
	}
	return fmt.Errorf("method not allowed! %s", r.Method)
}
