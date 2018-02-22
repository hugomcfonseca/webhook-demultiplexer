package apis

import (
	"fmt"

	"github.com/hugomcfonseca/cachet"
)

// CachetConfig ...
type CachetConfig struct {
	cachetURL string
	authType  string
	username  string
	password  string
	token     string
}

var cachetClient cachet.Client

// initClient set up a client to connect to Cachet API
func initClient() {
	cachetConfig := CachetConfig{}

	cachetClient, err := cachet.NewClient(cachetConfig.cachetURL, nil)

	if err != nil {
		fmt.Printf("Error creating Cachet client: %s", err)
		return
	}

	_, resp, err := cachetClient.General.Ping()

	if resp.StatusCode != 200 {
		fmt.Printf("Cachet server is not reachable: %s", err)
		return
	}
}

// getComponentByTag ...
func getComponentByTag(tag string) (bool, string) {
	var errorMsg string

	components, _, err := cachetClient.Components.GetAll()

	if err != nil {
		return false, ""
	}

	if components.Meta.Pagination.Count == 0 {
		errorMsg = fmt.Sprintf("Not found a component to tag '%s'\n", tag)
		return false, errorMsg
	}

	// find component name and id

	return true, ""
}

// getComponentStatus ...
func getComponentStatus(cid int) (bool, int) {
	component, _, err := cachetClient.Components.Get(cid)

	if err != nil {
		return false, cachet.ComponentStatusUnknown
	}

	if component.Status != cachet.ComponentStatusOperational {
		return false, component.Status
	}

	return true, component.Status
}

// reportIncident ...
func reportIncident(name string, msg string, status int, cID int, cStatus int) bool {
	component := &cachet.Incident{
		Name:            name,
		Message:         msg,
		Status:          status,
		Visible:         cachet.IncidentVisibilityPublic,
		ComponentID:     cID,
		ComponentStatus: cStatus,
	}

	incident, _, err := cachetClient.Incidents.Create(component)

	if err != nil || incident.ID <= 0 {
		return false
	}

	return true
}

// getIncidentByComponent gets all 
func getIncidentByComponent(cid int) (bool, int) {
	incidents, _, err := cachetClient.Incidents.GetAll()

	if err != nil {
		return false, -1
	}

	if incidents.Meta.Pagination.Count == 0 {
		//errorMsg = fmt.Sprintf("Not found an incident to component '%d'\n", cid)
		return false, -1
	}

	// find incident ID

	return true, 0
}

// updateIncident ...
func updateIncident() bool {

	return true
}
