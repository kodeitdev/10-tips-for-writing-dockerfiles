package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Handler handles request to endpoint /echo
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Received request on: %s%s ", r.Host, r.URL.Path)

	response := fmt.Sprint("My response: pong")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(response))
	if err != nil {
		log.Errorf("Could not send response because of an error: %s", err.Error())
	}
}

func main() {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		log.SetLevel(log.DebugLevel)
	}

	http.HandleFunc("/ping", Handler)

	listenAddress := "0.0.0.0:8080"

	log.Infof("Starting http server at address: %s", listenAddress)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatalf("HTTP server returned with error: %s", err.Error())
	}
}
