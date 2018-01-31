package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/hugomcfonseca/cachet"
)

const (
	version = "0.0.1"
)

var (
	ip   = flag.String("ip", "0.0.0.0", "Domain or IP address where hookserver is listen")
	port = flag.Int("port", 8080, "Port where hookserver is listen")
)

func main() {
	var webhookURL string

	router := mux.NewRouter()
	webhookURL = "http://localhost:8080/prefix/{id}"

	router.HandleFunc(webhookURL, hookHandler)

}

func hookHandler(w http.ResponseWriter, r *http.Request) {

}
