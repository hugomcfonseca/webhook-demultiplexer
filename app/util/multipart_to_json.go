package util

import (
	"net/http"
	"strconv"
)

// Tag ...
type Tag struct {
	Data string
}

// StatusCake ...
type StatusCake struct {
	Token      string
	Status     bool
	StatusCode int
	Site       string
	IP         string
	Tags       []Tag
	TestName   string
	CheckRate  string
}

func multipartToJSON(req *http.Request) (*StatusCake, error) {
	status := false

	req.ParseForm()
	statusCode, err := strconv.Atoi(req.FormValue("StatusCode"))

	if req.FormValue("Status") == "Up" {
		status = true
	}

	respJSON := &StatusCake{
		Token:      req.FormValue("Token"),
		Status:     status,
		StatusCode: statusCode,
		Site:       req.FormValue("URL"),
		IP:         req.FormValue("IP"),
		//Tags:       req.FormValue("Tags"),
		TestName:  req.FormValue("Name"),
		CheckRate: req.FormValue("CheckRate"),
	}

	return respJSON, err
}
