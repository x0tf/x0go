package schema

// Namespace represents a namespace
type Namespace struct {
	ID     string `json:"id"`
	Token  string `json:"token"`
	Active bool   `json:"active"`
}

// Namespaces represents a paginated namespace response
type Namespaces struct {
	Data     []*Namespace        `json:"data"`
	Metadata *PaginationMetadata `json:"pagination"`
}

// CreateNamespaceData represents the request model required for a namespace creation request
type CreateNamespaceData struct {
	ID         string `json:"id"`
	InviteCode string `json:"invite_code,omitempty"`
}

// PatchNamespaceData represents the request model required for a namespace patching request
type PatchNamespaceData struct {
	Active *bool `json:"active,omitempty"`
}
