package schema

// PaginationMetadata represents the meta data of a paginated response
type PaginationMetadata struct {
	TotalElements     int `json:"total_elements"`
	DisplayedElements int `json:"displayed_elements"`
}
