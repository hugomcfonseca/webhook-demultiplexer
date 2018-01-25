package cachet

import (
	_ "encoding/json"
	"fmt"
	_ "net/http"

	"github.com/hugomcfonseca/cachet"
)

func main() {
	client, err := cachet.NewClient("https://cachet.localhost", nil) // provide URL from cmdline

	if err != nil {
		fmt.Printf("Error creating Cachet client: %s", err)
		return
	}

	_, resp, err := client.General.Ping()

	if resp.Status != "200" {
		fmt.Printf("Cachet server is not reachable: %s", err)
		return
	}
}

func 