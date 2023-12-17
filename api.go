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
	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountByID), s.store))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// NewAPIserver creates a new instance of the APIserver.
func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}

//routes
func (a *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
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
