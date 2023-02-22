package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Structure reflecting API response results.
type ListMembersResponse struct {

	// Top layer of json struct in the response body.
	Members []Member `json:"members"`
}

// Lists members for a specific group.
func (a *Api) ListMembers(gn string) ([]Member, error) {

	query := url.Values{}
	query.Set("key", a.config.Account.APIKey)
	query.Set("gn", gn)

	URL := fmt.Sprintf("%s/members?%s", a.config.Account.BaseURL, query.Encode())

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

	var membersResponse ListMembersResponse
	if err := json.Unmarshal(body, &membersResponse); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %w", err)
	}

	return membersResponse.Members, nil
}
