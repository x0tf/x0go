package x0

import "github.com/valyala/fasthttp"

// Client represents the entrypoint to use the x0 API client
type Client struct {
	Instance   *InstanceClient
	Namespaces *NamespacesClient
}

// NewClient creates a new x0 API client
func NewClient(instance string) *Client {
	return NewClientWithToken(instance, "")
}

// NewClientWithToken creates a new x0 API client with an auth token set
func NewClientWithToken(instance, token string) *Client {
	http := &httpClient{
		instance: instance,
		token:    token,
		http: &fasthttp.Client{
			Name: "x0tf/x0go",
		},
	}

	return &Client{
		Instance:   &InstanceClient{http: http},
		Namespaces: &NamespacesClient{http: http},
	}
}
