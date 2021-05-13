package x0

import (
	"encoding/json"
	"fmt"

	"github.com/x0tf/x0go/schema"
)

// NamespacesClient is used to execute API requests directed to the namespace-related scopes
type NamespacesClient struct {
	http *httpClient
}

// GetList requests a paginated list of existent namespaces
func (client *NamespacesClient) GetList(limit, skip int) (*schema.Namespaces, error) {
	if limit <= 0 {
		limit = 10
	}
	if skip < 0 {
		skip = 0
	}

	response, err := client.http.execute("GET", fmt.Sprintf("%s?limit=%d&skip=%d", endpointNamespaces, limit, skip), nil)
	if err != nil {
		return nil, err
	}

	namespaces := new(schema.Namespaces)
	if err := json.Unmarshal(response, namespaces); err != nil {
		return nil, err
	}
	return namespaces, nil
}

// GetSingle requests a single namespace
func (client *NamespacesClient) GetSingle(id string) (*schema.Namespace, error) {
	response, err := client.http.execute("GET", fmt.Sprintf("%s/%s", endpointNamespaces, id), nil)
	if err != nil {
		return nil, err
	}

	namespace := new(schema.Namespace)
	if err := json.Unmarshal(response, namespace); err != nil {
		return nil, err
	}
	return namespace, nil
}

// Create creates a new namespace
func (client *NamespacesClient) Create(data *schema.CreateNamespaceData) (*schema.Namespace, error) {
	response, err := client.http.execute("POST", endpointNamespaces, data)
	if err != nil {
		return nil, err
	}

	namespace := new(schema.Namespace)
	if err := json.Unmarshal(response, namespace); err != nil {
		return nil, err
	}
	return namespace, nil
}

// Patch patches an existing namespace
func (client *NamespacesClient) Patch(id string, data *schema.CreateNamespaceData) (*schema.Namespace, error) {
	response, err := client.http.execute("PATCH", fmt.Sprintf("%s/%s", endpointNamespaces, id), data)
	if err != nil {
		return nil, err
	}

	namespace := new(schema.Namespace)
	if err := json.Unmarshal(response, namespace); err != nil {
		return nil, err
	}
	return namespace, nil
}

// ResetToken resets the token of an existing namespace
func (client *NamespacesClient) ResetToken(id string) (*schema.Namespace, error) {
	response, err := client.http.execute("POST", fmt.Sprintf("%s/%s/%s", endpointNamespaces, id, endpointPartNamespacesResetToken), nil)
	if err != nil {
		return nil, err
	}

	namespace := new(schema.Namespace)
	if err := json.Unmarshal(response, namespace); err != nil {
		return nil, err
	}
	return namespace, nil
}

// Delete deletes an existing namespace
func (client *NamespacesClient) Delete(id string) error {
	_, err := client.http.execute("DELETE", fmt.Sprintf("%s/%s", endpointNamespaces, id), nil)
	return err
}
