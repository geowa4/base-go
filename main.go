package main

import (
	"log"
	"net/http"
	"os"

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

func main() {
	rootMux := http.NewServeMux()
	rootMux.Handle("/", static.NewStaticMux())
	addr := getAddr(os.Getenv("GOPORT"))
	log.Fatal(http.ListenAndServe(addr, rootMux))
}
