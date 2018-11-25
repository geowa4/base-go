package webui

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/geowa4/base-go/components/webui/internal/assets"
	"github.com/rs/zerolog/log"
)

//go:generate go-bindata -pkg assets -o internal/assets/embeds.go html/...

type indexData struct {
	Greeting string
}

func index(w http.ResponseWriter, r *http.Request) {
	contents, err := assets.Asset("html/index.html")
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving index.html contents.")
	}
	t := template.Must(template.New("index").Parse(string(contents)))
	var b strings.Builder
	err = t.Execute(&b, indexData{
		Greeting: "Hello world!",
	})
	if err == nil {
		_, err := w.Write([]byte(b.String()))
		if err != nil {
			log.Error().Err(err).Msg("Error writing index.html to response.")
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Error rendering index template.")
	}
}
