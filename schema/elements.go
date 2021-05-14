package schema

// ElementType represents the type an element may have
type ElementType string

const (
	ElementTypePaste    = ElementType("PASTE")
	ElementTypeRedirect = ElementType("REDIRECT")
)

// Element represents an element
type Element struct {
	Namespace  string                 `json:"namespace"`
	Key        string                 `json:"key"`
	Type       ElementType            `json:"type"`
	Data       map[string]interface{} `json:"public_data"`
	Views      int                    `json:"views"`
	MaxViews   int                    `json:"max_views"`
	ValidFrom  int64                  `json:"valid_from"`
	ValidUntil int64                  `json:"valid_until"`
	Created    int64                  `json:"created"`
}

// Elements represents a paginated element response
type Elements struct {
	Data     []*Element          `json:"data"`
	Metadata *PaginationMetadata `json:"pagination"`
}

// CreateBaseElementData represents the basic data every element creation request body must have
type CreateBaseElementData struct {
	Key        *string `json:"key,omitempty"`
	MaxViews   *int    `json:"max_views,omitempty"`
	ValidFrom  *int64  `json:"valid_from,omitempty"`
	ValidUntil *int64  `json:"valid_until,omitempty"`
}

// CreatePasteElementData represents the request model required for a paste element creation request
type CreatePasteElementData struct {
	Content string `json:"content"`
	CreateBaseElementData
}

// CreateRedirectElementData represents the request model required for a redirect element creation request
type CreateRedirectElementData struct {
	TargetURL string `json:"target_url"`
	CreateBaseElementData
}

// PatchElementData represents the request model required for a general element patching request
type PatchElementData CreateBaseElementData
