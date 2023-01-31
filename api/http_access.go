package api

import (
	"fmt"
	"net/http"
)

// Custom HTTP call.
// If failed, retry a default number of times.
func (a *Api) httpCall(url string) (*http.Response, error) {

	retries := 3

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequest: %w", err)
	}

	var resp *http.Response
	success := false
	for retries > 0 {
		// > If the returned error is nil, the Response will contain a non-nil
		// > Body which the user is expected to close.
		resp, err = client.Do(req)

		if err != nil {
			// Invalid URL (Like different scheme) etc.
			retries -= 1
			continue
		}

		if resp.StatusCode == http.StatusOK {
			// Success!
			success = true
			break
		}

		if resp.StatusCode/100 == 4 {
			// If the StatusCode starts with 4, it is user's error,
			// so it should not be retried.
			return nil, fmt.Errorf("failed to client.Do: StatusCode is %d", resp.StatusCode)
		}

		retries -= 1
	}

	if !success {
		return nil, fmt.Errorf("failed to client.Do after several retries.")
	}

	return resp, nil
}
