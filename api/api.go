package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	listnAddr string
}

type ApiError struct {
	Err string
}

func NewAPIServer(listnAddr string) *APIServer {
	return &APIServer{
		listnAddr: listnAddr,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()
	router.HandleFunc("/account", makeHTTPHandler(s.handleAccount))

	log.Println("Starting HTTP listener at", s.listnAddr)

	http.ListenAndServe(s.listnAddr, router)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err == nil {
			WriteJson(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		return s.handleGetAccount(w, r)
	case r.Method == "POST":
		return s.handleCreateAccount(w, r)
	case r.Method == "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method not allowed: %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
