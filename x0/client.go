package x0

import "github.com/valyala/fasthttp"

// Client represents the entrypoint to use the x0 API client
type Client struct {
	Instance *InstanceClient
}

// NewClient creates a new x0 API client
func NewClient(instance string) *Client {
	http := &httpClient{
		instance: instance,
		http: &fasthttp.Client{
			Name: "x0tf/x0go",
		},
	}

	return &Client{
		Instance: &InstanceClient{http: http},
	}
}
