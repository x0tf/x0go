package x0

import (
	"encoding/json"
	"fmt"

	"github.com/x0tf/x0go/schema"
)

// InvitesClient is used to execute API requests directed to the invite-related scopes
type InvitesClient struct {
	http *httpClient
}

// GetList requests a paginated list of existent invites
func (client *InvitesClient) GetList(limit, skip int) (*schema.Invites, error) {
	if limit <= 0 {
		limit = 10
	}
	if skip < 0 {
		skip = 0
	}

	response, err := client.http.execute("GET", fmt.Sprintf("%s?limit=%d&skip=%d", endpointInvites, limit, skip), nil)
	if err != nil {
		return nil, err
	}

	invites := new(schema.Invites)
	if err := json.Unmarshal(response, invites); err != nil {
		return nil, err
	}
	return invites, nil
}

// GetSingle requests a single invite
func (client *InvitesClient) GetSingle(code string) (*schema.Invite, error) {
	response, err := client.http.execute("GET", fmt.Sprintf("%s/%s", endpointInvites, code), nil)
	if err != nil {
		return nil, err
	}

	invite := new(schema.Invite)
	if err := json.Unmarshal(response, invite); err != nil {
		return nil, err
	}
	return invite, nil
}

// Create creates a new invite
func (client *InvitesClient) Create(data *schema.CreateInviteData) (*schema.Invite, error) {
	response, err := client.http.execute("POST", endpointInvites, data)
	if err != nil {
		return nil, err
	}

	invite := new(schema.Invite)
	if err := json.Unmarshal(response, invite); err != nil {
		return nil, err
	}
	return invite, nil
}

// Patch patches an existing invite
func (client *InvitesClient) Patch(code string, data *schema.PatchInviteData) (*schema.Invite, error) {
	response, err := client.http.execute("PATCH", fmt.Sprintf("%s/%s", endpointInvites, code), data)
	if err != nil {
		return nil, err
	}

	invite := new(schema.Invite)
	if err := json.Unmarshal(response, invite); err != nil {
		return nil, err
	}
	return invite, nil
}

// Delete deletes an existing invite
func (client *InvitesClient) Delete(code string) error {
	_, err := client.http.execute("DELETE", fmt.Sprintf("%s/%s", endpointInvites, code), nil)
	return err
}
