package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/hugomcfonseca/cachet"
)

const (
	version = "1"
)

var (
	auth    = flag.String("auth", "", "Authentication token that StatusCake webhook POSTs must match")
	configs = flag.String("configs", "configs.json", "JSON file with configurations to target APIs")
	address = flag.String("url", "0.0.0.0", "Domain or IP address where hookserver is listen")
	port    = flag.Int("port", 8080, "Port where hookserver is listen")
	prefix  = flag.String("prefix", "v1", "")
)

// Config set an object containing configurations by each destination API
type Config struct {
	URL      string `json:"url"`
	AuthType string `json:"auth_type"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func main() {
	flag.Parse()

	isValid, webhookURL := verifyArgs()
	if !isValid {
		log.Printf("%s", webhookURL)
		return
	}

	router := mux.NewRouter()
	router.PathPrefix(*prefix)
	router.HandleFunc("/", listTargetsHandler).Methods("GET")
	router.HandleFunc("/{application}", sendNotifyHandler).Methods("POST")
	//router.Use(loggerMiddleware, authorizationMiddleware)
	//http.Handle(webhookURL, loggerMiddleware(authorizationMiddleware(router)))
	log.Printf("%s", webhookURL)
	log.Fatal(http.ListenAndServe(webhookURL, router))
}

// verifyArgs performs arguments validation and return webhook URL on success
func verifyArgs() (bool, string) {
	var errMsg string
	var webhookURL string

	if *auth == "" {
		errMsg = "Invalid authentication! It must not be empty.\n"
		return false, errMsg
	}

	if *address == "" {
		errMsg = "Invalid URL! It must not be empty.\n"
		return false, errMsg
	}

	if *port < 1025 || *port > 65535 {
		errMsg = "Invalid port! Please, check it is between 1025 and 65535.\n"
		return false, errMsg
	}

	if *prefix != "" {
		*prefix = strings.Trim(*prefix, "/")
	}

	webhookURL = fmt.Sprintf("%s:%d", *address, *port)

	return true, webhookURL
}

// loggerMiddleware Middleware level to log API requests
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("[%s]\t%s\t%s", r.Method, r.URL.String(), time.Since(start))
	})
}

// authorizationMiddleware ...
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			vars := mux.Vars(r)

			if strings.Compare(vars["webhook_token"], *auth) != 0 {
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// listTargetsHandler Handler to redirect request to specified application
func listTargetsHandler(w http.ResponseWriter, r *http.Request) {

}

// sendNotifyHandler Handler to redirect request to specified application
func sendNotifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Printf("Received hook request to %v with token %v\n", vars["application"])
}

// loadAppConfigs ...
func loadAppConfigs(app string) ([]Config, error) {
	var config []Config

	raw, err := ioutil.ReadFile(*configs)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	json.Unmarshal(raw, &config)

	return config, err
}
