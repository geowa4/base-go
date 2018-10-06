package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/geowa4/base-go/components/api"
	"github.com/geowa4/base-go/components/webui"
	"github.com/justinas/alice"
	"github.com/rs/zerolog/log"
)

func getAppPort() uint16 {
	intPort, err := strconv.ParseUint(os.Getenv("GO_APP_PORT"), 10, 16)
	if err != nil {
		return 8000
	}
	return uint16(intPort)
}

func newAppMux(accessLogChain alice.Chain) *http.ServeMux {
	hlogChain := newAccessLogChain(log.Logger)
	mux := http.NewServeMux()
	mux.Handle("/graphql", hlogChain.Then(api.NewGraphQLHandler()))
	mux.Handle("/", hlogChain.Then(webui.NewStaticMux()))
	return mux
}

func startAppServer(ctx context.Context, config *webServerConfig) {
	appAddr := fmt.Sprintf(":%d", getAppPort())
	appMux := newAppMux(config.accessLogChain)
	startServer(ctx, appAddr, appMux, config.idleAppConnectionsClosed)
}
