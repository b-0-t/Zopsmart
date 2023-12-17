package main

import "net/http"

// APIserver represents the configuration for the API server.
type APIserver struct {
	listenAddr string
}

// NewAPIserver creates a new instance of the APIserver.
func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}

//routes
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
