package apis

import (
	"fmt"

	"github.com/hugomcfonseca/cachet"
)

// CachetConfig returns configurations used for creating Cachet client
type CachetConfig struct {
	CachetURL string ``
	AuthType  string ``
	Username  string ``
	Password  string ``
	Token     string ``
}

// CachetClient returns a pointer to Cachet client
type CachetClient struct {
	client *cachet.Client
}

// InitClient set up a client to connect to Cachet API
func InitClient(config *CachetConfig) *cachet.Client {
	client, err := cachet.NewClient(config.CachetURL, nil)

	if err != nil {
		fmt.Printf("Error creating Cachet client: %s", err)
		return nil
	}

	_, resp, err := client.General.Ping()

	if resp.StatusCode != 200 || err == nil {
		fmt.Printf("Cachet server is not reachable: %s", err)
		return nil
	}

	switch config.AuthType {
	case "token":
		client.Authentication.SetTokenAuth(config.Token)
	case "basic":
		client.Authentication.SetBasicAuth(config.Username, config.Password)
	default:
		return nil
	}

	return client
}

// GetComponentByTag returns all existing components that are enabled and matches
// with returned tags from StatusCake
func (s *CachetClient) GetComponentByTag(tag string) (bool, *cachet.Component, string) {
	var res *cachet.Component
	var errMsg string
	isEvaluable := true

	filters := &cachet.ComponentsQueryParams{
		Name:    tag,
		Enabled: true,
	}
	components, _, err := s.client.Components.GetAll(filters)

	if err != nil {
		return false, res, ""
	}

	switch {
	case components.Meta.Pagination.Count <= 0:
		errMsg = fmt.Sprintf("Not found a component to tag '%s'\n", tag)
		isEvaluable = false
	case components.Meta.Pagination.Count == 1:
		errMsg = fmt.Sprintf("Fabulous! Found it a component matching tag '%s'\n", tag)
	case components.Meta.Pagination.Count > 1:
		errMsg = fmt.Sprintf("Found multiple components matching tag '%s'. Please, refine your tags.\n", tag)
		isEvaluable = false
	default:
		errMsg = fmt.Sprint("Unknown error reading returned results...")
		isEvaluable = false
	}

	if !isEvaluable {
		return isEvaluable, res, errMsg
	}

	for _, component := range components.Components {
		res = &cachet.Component{
			ID:     component.ID,
			Name:   component.Name,
			Status: component.Status,
		}
	}

	return true, res, errMsg
}

// GetIncidentByComponent gets all incidents by a given component
// @todo: using NOT notation in querystring fields
func (s *CachetClient) GetIncidentByComponent(cid int, cName string) (bool, *cachet.Incident, string) {
	var res *cachet.Incident
	var errMsg string
	isEvaluable := true

	filters := &cachet.IncidentsQueryParams{
		ComponentID: cid,
	}
	incidents, _, err := s.client.Incidents.GetAll(filters)

	if err != nil {
		return false, res, errMsg
	}

	switch {
	case incidents.Meta.Pagination.Count <= 0:
		errMsg = fmt.Sprintf("Not found a component to tag '%s'\n", cName)
		isEvaluable = false
	case incidents.Meta.Pagination.Count == 1:
		errMsg = fmt.Sprintf("Fabulous! Found it a component matching tag '%s'\n", cName)
	case incidents.Meta.Pagination.Count > 1:
		errMsg = fmt.Sprintf("Found multiple components matching tag '%s'. Please, refine your tags.\n", cName)
		isEvaluable = false
	default:
		errMsg = fmt.Sprint("Unknown error reading returned results...")
		isEvaluable = false
	}

	if !isEvaluable {
		return isEvaluable, res, errMsg
	}

	for _, incident := range incidents.Incidents {
		if incident.Status != cachet.IncidentStatusFixed {
			res = &incident
		}
	}

	return true, res, ""
}

// ReportIncident creates an incident in case no incident was opened
// and it is related to the target component/tag
func (s *CachetClient) ReportIncident(incident *cachet.Incident) (bool, error) {
	incident, _, err := s.client.Incidents.Create(incident)

	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateIncident creates an incident update in case it is already opened
// an incident related to the target component/tag
//
// Uses cases:
//  - creates an incident update with (or without) updating incident status
//  - creates an incident update that changes status to resolved and fix the incident
func (s *CachetClient) UpdateIncident(incident *cachet.Incident, status bool) (bool, error) {
	update := &cachet.IncidentUpdate{
		IncidentID:      incident.ID,
		ComponentID:     incident.ComponentID,
		ComponentStatus: 1, // fixes me
		Status:          1, // fixes me
		Message:         "",
	}

	_, _, err := s.client.IncidentUpdates.Create(incident.ID, update)

	if err != nil {
		return false, err
	}

	return true, nil
}
