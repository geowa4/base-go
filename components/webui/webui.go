package webui

import "net/http"

// NewStaticMux Does stuff
func NewStaticMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}
