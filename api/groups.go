package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Structure reflecting API response results.
type ListGroupsResponse struct {

	// Top layer of json struct in the response body.
	Groups []Group `json:"groups"`
}

// Lists of all groups.
// Fetch data from {BASE_URL}/groups with api_key as a query parameter.
func (a *Api) ListGroups() ([]Group, error) {

	var (
		URL = fmt.Sprintf("%s/groups?key=%s", a.config.Account.BaseURL, a.config.Account.APIKey)
	)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to http.NewRequest: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to client.Do: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to io.ReadAll: %w", err)
	}

	var groupsResponse ListGroupsResponse
	if err := json.Unmarshal(body, &groupsResponse); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %w", err)
	}

	return groupsResponse.Groups, nil
}
