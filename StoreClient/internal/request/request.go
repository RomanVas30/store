package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonRequest(method RequestMethod, url string, body interface{}, token string) (*http.Response, error) {
	bytesRepresentation, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
