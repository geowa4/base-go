package static

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// NewStaticMux Does stuff
func NewStaticMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}
