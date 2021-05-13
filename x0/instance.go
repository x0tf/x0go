package x0

import (
	"encoding/json"

	"github.com/x0tf/x0go/schema"
)

// InstanceClient is used to execute API requests directed to the general instance
type InstanceClient struct {
	http *httpClient
}

// GetInfo requests the instance information for the specified x0 instance
func (client *InstanceClient) GetInfo() (*schema.InstanceInfo, error) {
	response, err := client.http.execute("GET", endpointInfo, nil)
	if err != nil {
		return nil, err
	}

	info := new(schema.InstanceInfo)
	if err := json.Unmarshal(response, info); err != nil {
		return nil, err
	}
	return info, nil
}
