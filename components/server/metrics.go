package server

import (
	"context"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/justinas/alice"
)

func getMetricsPort() uint16 {
	intPort, err := strconv.ParseUint(os.Getenv("GO_METRICS_PORT"), 10, 16)
	if err != nil {
		return getAppPort() + 1
	}
	return uint16(intPort)
}

func newMetricsMux(accessLogChain alice.Chain) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", accessLogChain.Then(expvar.Handler()))
	return mux
}
func startMetricsServer(ctx context.Context, config *webServerConfig) {
	metricsAddr := fmt.Sprintf(":%d", getMetricsPort())
	metricsMux := newMetricsMux(config.accessLogChain)
	startServer(ctx, metricsAddr, metricsMux, config.idleMetricsConnectionsClosed)
}
