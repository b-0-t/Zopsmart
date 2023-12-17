package main

import {
	"net/http"
	"github.com/gorilla/mux"
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}


type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// APIserver represents the configuration for the API server.
type APIserver struct {
	listenAddr string
}

func (a *APIserver) Run() {
	router:=mux.NewRouter()
	router.HandleFunc("/account",a.handleAccount)
}

// NewAPIserver creates a new instance of the APIserver.
func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}

//routes
func (a *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}


func (a *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *APIserver) handleTansfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
