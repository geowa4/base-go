package main

import (
	"expvar"
	"net/http"
	"os"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"github.com/geowa4/base-go/components/api"
	"github.com/geowa4/base-go/components/static"
)

func getAddr(port string) (addr string) {
	addr = ":"
	if port == "" {
		addr += "8000"
	} else {
		addr += port
	}
	return
}

func configureGlobalLogging(level string) {
	zerolog.TimeFieldFormat = ""
	if logLevel, err := zerolog.ParseLevel(level); err != nil {
		zerolog.SetGlobalLevel(logLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func makeAccessLogChain(log zerolog.Logger) (chain alice.Chain) {
	chain = alice.New()

	// Install the logger handler with default output on the console
	chain = chain.Append(hlog.NewHandler(log))

	// Install some provided extra handler to set some request's context fields.
	// Thanks to those handler, all our logs will come with some pre-populated fields.
	chain = chain.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	chain = chain.Append(hlog.RemoteAddrHandler("ip"))
	chain = chain.Append(hlog.UserAgentHandler("user_agent"))
	chain = chain.Append(hlog.RefererHandler("referer"))
	chain = chain.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
	return
}

func startServer(mux *http.ServeMux) {
	addr := getAddr(os.Getenv("GOPORT"))
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal().Err(err).Msg("Startup failed.")
	}
}

func main() {
	configureGlobalLogging(os.Getenv("GOLOGLEVEL"))
	hlogChain := makeAccessLogChain(log.Logger)
	rootMux := http.NewServeMux()
	rootMux.Handle("/debug/vars", hlogChain.Then(expvar.Handler()))
	rootMux.Handle("/graphql", hlogChain.Then(api.NewGraphQLHandler()))
	rootMux.Handle("/", hlogChain.Then(static.NewStaticMux()))
	startServer(rootMux)
}
