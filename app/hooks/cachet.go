package cachet

import (
	_ "encoding/json"
	"fmt"
	_ "net/http"

	"github.com/hugomcfonseca/cachet"
)

type CachetConfig struct {
	cachetURL string
	authType  string
	username  string
	password  string
	token     string
}

var cachetClient cachet.Client

func cachetMain() {
	cachetConfig := CachetConfig{}

	cachetClient, err := cachet.NewClient(cachetConfig.cachetURL, nil)

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

func tagToComponent(tag string) (string, bool) {
	componentResp, resp, err := cachetClient.Components.GetAll()

	if err != nil {
		return "", false
	}

	return "", true
}

func checkComponentStatus(compID int) (int, bool) {

	return 0, true
}

func reportIncident(name string, msg string, status int, compID int, compStatus int) bool {
	component := &cachet.Incident{
		Name:            name,
		Message:         msg,
		Status:          cachet.ComponentStatusOperational, // fixes me
		Visible:         cachet.IncidentVisibilityPublic,
		ComponentID:     compID,
		ComponentStatus: cachet.ComponentStatusPartialOutage, // fixes me
	}

	incidentResp, resp, err := cachetClient.Incidents.Create(component)

	return true
}

func updateIncident() {

}
