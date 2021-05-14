package schema

// Invite represents an invite
type Invite struct {
	Code    string `json:"code"`
	Uses    int    `json:"uses"`
	MaxUses int    `json:"max_uses"`
	Created int64  `json:"created"`
}

// Invites represents a paginated invite response
type Invites struct {
	Data     []*Invite           `json:"data"`
	Metadata *PaginationMetadata `json:"pagination"`
}

// BaseInviteData represents the request model both the invite creation and invite patching endpoints require
type BaseInviteData struct {
	Code    *string `json:"code,omitempty"`
	MaxUses *int    `json:"max_uses,omitempty"`
}

// CreateInviteData represents the request model the invite creation endpoint requires
// We are simply aliasing the BaseInviteData structure here to allow later additions
type CreateInviteData BaseInviteData

// PatchInviteData represents the request model the invite patching endpoint requires
// We are simply aliasing the BaseInviteData structure here to allow later additions
type PatchInviteData BaseInviteData
