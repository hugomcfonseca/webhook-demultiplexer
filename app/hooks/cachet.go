package cachet

import (
	_ "encoding/json"
	"fmt"
	_ "net/http"

	"github.com/hugomcfonseca/cachet"
)

var cachetClient cachet.Client

func main() {
	cachetClient, err := cachet.NewClient("https://cachet.localhost", nil) // provide URL from cmdline

	if err != nil {
		fmt.Printf("Error creating Cachet client: %s", err)
		return
	}

	_, resp, err := cachetClient.General.Ping()

	if resp.Status != "200" {
		fmt.Printf("Cachet server is not reachable: %s", err)
		return
	}
}

func tagsToComponents() (string, bool) {
	componentResp, resp, err := cachetClient.Components.GetAll()

	if err != nil {
		return "", false
	}

	return "", true
}

func reportIncident(name string, description string, status int, compId string, compStatus int) {
	component := cachet.Incident{
		Name:        name,
		Description: description,
		Status:      cachet.ComponentStatusOperational, // fixes me
		Visible: 	 cachet.IncidentVisibilityPublic,
		ComponentID: compId,
		ComponentStatus: cachet.ComponentStatusPartialOutage // fixes me
	}

	incidentResp, resp, err := cachetClient.Incidents.Create()
}
