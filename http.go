package simgo5

import (
	"io"
	"net/http"
)

// newHttpClient creates and returns a new HTTP client instance
func newHttpClient() *http.Client {
	return &http.Client{}
}

// executeRequest performs an HTTP GET request to the specified URL with the given API key
// It returns the response body as a byte slice or an error if the request fails
func executeRequest(client *http.Client, url, apiKey string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, httpError
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
