package x0

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	ErrorCodeGenericUnexpectedError   = 1000
	ErrorCodeGenericUnauthorized      = 1001
	ErrorCodeGenericBadQueryParameter = 1002
	ErrorCodeGenericBadRequestBody    = 1003
	ErrorCodeGenericNamespaceNotFound = 1004
	ErrorCodeGenericElementNotFound   = 1005
	ErrorCodeGenericInviteNotFound    = 1006

	ErrorCodeNamespaceIllegalNamespaceID = 2000
	ErrorCodeNamespaceNamespaceIDInUse   = 2001
	ErrorCodeNamespaceInvalidInviteCode  = 2002

	ErrorCodeElementElementKeyInUse          = 3000
	ErrorCodeElementNamespaceDeactivated     = 3001
	ErrorCodeElementPasteEmptyPasteContent   = 3100
	ErrorCodeElementRedirectInvalidTargetURL = 3200

	ErrorCodeInviteInvitesDisabled = 4000
	ErrorCodeInviteInviteCodeInUse = 4001
)

// APIError represents a x0-specific error structure returned from the x0 API
type APIError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Error formats a x0-specific error
func (err *APIError) Error() string {
	return fmt.Sprintf("%d: %s (custom_data: %v)", err.Code, err.Message, err.Data)
}

func errFromResponse(response *fasthttp.Response) error {
	status := response.StatusCode()
	body := response.Body()

	if status < 200 || status > 299 {
		apiErr := new(APIError)
		if err := json.Unmarshal(body, apiErr); err == nil {
			return apiErr
		}

		return errors.New(string(body))
	}

	return nil
}
