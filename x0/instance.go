package x0

import "encoding/json"

// InstanceInfo represents the application information about a x0 instance
type InstanceInfo struct {
	Version    string            `json:"version"`
	Production bool              `json:"production"`
	Settings   *InstanceSettings `json:"settings"`
}

// InstanceSettings represents the settings of a x0 instance
type InstanceSettings struct {
	InvitesEnabled   bool              `json:"invites"`
	NamespaceIDRules *NamespaceIDRules `json:"namespace_id_rules"`
}

// NamespaceIDRules represents the rules a namespace ID has to match for a x0 instance
type NamespaceIDRules struct {
	MinLength         int    `json:"min_length"`
	MaxLength         int    `json:"max_length"`
	AllowedCharacters string `json:"allowed_characters"`
}

// InstanceClient is used to execute API requests directed to the general instance
type InstanceClient struct {
	http *httpClient
}

// GetInfo requests the instance information for the specified instance
func (client *InstanceClient) GetInfo() (*InstanceInfo, error) {
	response, err := client.http.execute("GET", endpointInfo, nil)
	if err != nil {
		return nil, err
	}

	info := new(InstanceInfo)
	if err := json.Unmarshal(response, info); err != nil {
		return nil, err
	}
	return info, nil
}
