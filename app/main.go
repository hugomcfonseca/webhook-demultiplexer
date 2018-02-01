package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

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

type Config struct {
	HookToken string `json:"hook_token"`
	AppToken  string `json:"app_token"`
}

func main() {
	router := mux.NewRouter()
	webhookURL := "http://localhost:8080/prefix/{application}/{webhook_token}"

	router.HandleFunc(webhookURL, hookHandler)
	router.Use(authorizationMiddleware)
}

// authorizationMiddleware ...
func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		validAuth := "SECRET"

		if strings.Compare(vars["webhook_token"], validAuth) != 0 {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// hookHandler Handler to redirect request to specified application
func hookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println("Received hook request to %v with token %v\n", vars["application"], vars["webhook_token"])
}

// loadAppConfigs ...
func loadAppConfigs(app string) (Config, bool) {
	var config Config
	var filename string
	err := false

	switch app {
	case "cachet":
		filename = "configs/cachet.json"
	default:
		err = true
	}

	configFile, loadErr := os.Open(filename)
	defer configFile.Close()

	if loadErr != nil {
		return config, true
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, err
}
