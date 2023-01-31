package api

import (
	"encoding/json"
	"fmt"
	"io"
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

	resp, err := a.httpCall(URL)
	if err != nil {
		return nil, fmt.Errorf("failed to httpCall: %w", err)
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
