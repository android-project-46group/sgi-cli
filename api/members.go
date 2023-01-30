package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type ListMembersResponse struct {
	Members []Member `json:"members"`
}

// Lists members for a specific group.
func (a *Api) ListMembers(gn string) ([]Member, error) {

	var (
		URL = fmt.Sprintf("%s/members?gn=%s&key=%s", a.config.Account.BaseURL, gn, a.config.Account.APIKey)
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

	var membersResponse ListMembersResponse
	if err := json.Unmarshal(body, &membersResponse); err != nil {
		return nil, fmt.Errorf("failed to json.Unmarshal: %w", err)
	}

	return membersResponse.Members, nil
}
