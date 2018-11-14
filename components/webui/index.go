package webui

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

//go:generate go-bindata -pkg $GOPACKAGE -o embeds.go html/...

type indexData struct {
	Greeting string
}

func index(w http.ResponseWriter, r *http.Request) {
	contents, err := Asset("html/index.html")
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving index.html contents.")
	}
	t := template.Must(template.New("index").Parse(string(contents)))
	var b strings.Builder
	err = t.Execute(&b, indexData{
		Greeting: "Hello world!",
	})
	if err == nil {
		w.Write([]byte(b.String()))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Error rendering index template.")
	}
}
