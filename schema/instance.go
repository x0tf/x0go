package schema

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
