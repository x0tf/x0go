package x0

import (
	"encoding/json"
	"fmt"

	"github.com/x0tf/x0go/schema"
)

// ElementsClient is used to execute API requests directed to the element-related scopes
type ElementsClient struct {
	http *httpClient
}

// GetList requests a paginated list of existent elements
func (client *ElementsClient) GetList(limit, skip int) (*schema.Elements, error) {
	return client.GetListInNamespace("", limit, skip)
}

// GetListInNamespace requests a paginated list of existent elements in a specific namespace
func (client *ElementsClient) GetListInNamespace(namespaceID string, limit, skip int) (*schema.Elements, error) {
	if limit <= 0 {
		limit = 10
	}
	if skip < 0 {
		skip = 0
	}

	response, err := client.http.execute("GET", fmt.Sprintf("%s/%s?limit=%d&skip=%d", endpointElements, namespaceID, limit, skip), nil)
	if err != nil {
		return nil, err
	}

	elements := new(schema.Elements)
	if err := json.Unmarshal(response, elements); err != nil {
		return nil, err
	}
	return elements, nil
}

// CreatePaste creates a new paste element
func (client *ElementsClient) CreatePaste(namespaceID string, data *schema.CreatePasteElementData) (*schema.Element, error) {
	response, err := client.http.execute("POST", fmt.Sprintf("%s/%s/%s", endpointElements, namespaceID, endpointPartElementsPaste), data)
	if err != nil {
		return nil, err
	}

	element := new(schema.Element)
	if err := json.Unmarshal(response, element); err != nil {
		return nil, err
	}
	return element, nil
}

// CreateRedirect creates a new redirect element
func (client *ElementsClient) CreateRedirect(namespaceID string, data *schema.CreateRedirectElementData) (*schema.Element, error) {
	response, err := client.http.execute("POST", fmt.Sprintf("%s/%s/%s", endpointElements, namespaceID, endpointPartElementsPaste), data)
	if err != nil {
		return nil, err
	}

	element := new(schema.Element)
	if err := json.Unmarshal(response, element); err != nil {
		return nil, err
	}
	return element, nil
}

// Patch patches an existing element
func (client *ElementsClient) Patch(namespaceID, key string, data *schema.PatchElementData) (*schema.Element, error) {
	response, err := client.http.execute("PATCH", fmt.Sprintf("%s/%s/%s", endpointElements, namespaceID, key), data)
	if err != nil {
		return nil, err
	}

	element := new(schema.Element)
	if err := json.Unmarshal(response, element); err != nil {
		return nil, err
	}
	return element, nil
}

// Delete deletes an existing element
func (client *ElementsClient) Delete(namespaceID, key string) error {
	_, err := client.http.execute("DELETE", fmt.Sprintf("%s/%s/%s", endpointElements, namespaceID, key), nil)
	return err
}
