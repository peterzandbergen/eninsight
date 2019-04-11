// Implements the service as a new Google Cloud Run function.
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/peterzandbergen/eninsight/adapters/rest"
)

const (
	EnvPort     = "PORT"
	DefaultPort = "8080"
)

func main() {
	port := os.Getenv(EnvPort)
	if len(port) == 0 {
		port = DefaultPort
	}
	addr := ":" + port

	h := &rest.Handler{}

	fmt.Printf("Listening on %s\n", addr)
	http.ListenAndServe(addr, h.Routes())
}
