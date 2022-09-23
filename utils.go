package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// generic http get function
func get[T any](url string) (jsonResponse T, err error) {

	API_TOKEN := os.Getenv("TODOIST_API_KEY")

	// get request
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"Host":          {"www.host.com"},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + API_TOKEN},
	}

	resp, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	// convert response to json
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(body), &jsonResponse)

	return
}
