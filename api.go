package main

// APIserver represents the configuration for the API server
type APIserver struct {
	listenAddr string
}

// NewAPIserver creates a new instance of the APIserver
func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}
