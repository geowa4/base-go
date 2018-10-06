package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func newAccessLogChain(log zerolog.Logger) (chain alice.Chain) {
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

func startServer(webCtx context.Context, addr string, mux *http.ServeMux, idleConnectionsClosed chan bool) {
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server startup failed.")
		} else {
			log.Info().Msg(fmt.Sprintf("Server on %s shutting down.", addr))
		}
	}()
	<-webCtx.Done()
	shutdownContext, cancelShutdown := context.WithTimeout(context.Background(), 30*time.Second)
	if shutdownErr := srv.Shutdown(shutdownContext); shutdownErr != nil {
		log.Error().Err(shutdownErr).Msg("Server shutdown error.")
	}
	cancelShutdown()
	close(idleConnectionsClosed)
}

type webServerConfig struct {
	idleAppConnectionsClosed     chan bool
	idleMetricsConnectionsClosed chan bool
	accessLogChain               alice.Chain
	cancelContextFunc            func()
}

func (config *webServerConfig) Cancel() {
	log.Info().Msg("Gracefully shutting down servers.")
	config.cancelContextFunc()
	log.Info().Msg("Waiting for metrics connections to close.")
	<-config.idleMetricsConnectionsClosed
	log.Info().Msg("Waiting for app connections to close.")
	<-config.idleAppConnectionsClosed
}

// NewWebServers spins up app and metrics servers and
// returns a function to cancel the web contexts,
// resulting in their graceful shutdown.
// The app server comprises the ui and api.
func NewWebServers(log zerolog.Logger) func() {
	webConfig := &webServerConfig{
		idleAppConnectionsClosed:     make(chan bool, 1),
		idleMetricsConnectionsClosed: make(chan bool, 1),
		accessLogChain:               newAccessLogChain(log),
	}
	webCtx, cancelWeb := context.WithCancel(context.Background())
	webConfig.cancelContextFunc = cancelWeb
	go startAppServer(webCtx, webConfig)
	go startMetricsServer(webCtx, webConfig)
	return webConfig.Cancel
}
