package x0

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

const (
	PublicInstance        = "https://api.x0.tf"
	PublicStagingInstance = "https://api.s.x0.tf"

	endpointInfo = "/v2/info"
)

type httpClient struct {
	instance string
	http     *fasthttp.Client
}

func (client *httpClient) execute(method, endpoint string, payload interface{}) ([]byte, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.SetMethod(method)
	request.SetRequestURI(client.instance + endpoint)

	if payload != nil {
		payloadJson, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		request.SetBody(payloadJson)
	}

	if err := client.http.Do(request, response); err != nil {
		return nil, err
	}

	if err := errFromResponse(response); err != nil {
		return nil, err
	}
	return response.Body(), nil
}
