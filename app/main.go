package main

import (
	"flag"
	"net/http"

	_ "github.com/gorilla/mux"
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

}

func hookHandler(w http.ResponseWriter, r *http.Request) {

}
