package services_uitls_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SendPostRequest[TRequestBody any, TResponseData any](url string, body *TRequestBody, headers *map[string]string) (*TResponseData, error) {
	var _body *bytes.Buffer = bytes.NewBuffer([]byte{})

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		_body = bytes.NewBuffer(bodyBytes)
	}
	req, err := http.NewRequest(http.MethodPost, url, _body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if headers != nil {
		for key, value := range *headers {
			if value != "" {
				req.Header.Add(key, value)
			}
		}
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !IsStatusCode2xx(resp.StatusCode) {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	var response *TResponseData
	err = json.Unmarshal(responseBody, &response)
	return response, err
}
