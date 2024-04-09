package api_errors

import "net/http"

var (
	ErrUnauthorizedAccess        = NewAPIError(401, "Unauthorized access")
	ErrForbiddenAccess           = NewAPIError(403, "Forbidden access")
	ErrInvalidToken              = NewAPIError(401, "Invalid token")
	ErrBadRequest                = NewAPIError(400, "Bad request")
	ErrInvalidUUID               = NewAPIError(http.StatusBadRequest, "Invalid UUID")
	ErrRecordNotFound            = NewAPIError(404, "Record not found")
	ErrInvalidUserNameOrPassword = NewAPIError(400, "Invalid username and password")
	ErrExtensionMismatch         = NewAPIError(400, "file extension not supported")
	ErrThumbExtensionMismatch    = NewAPIError(400, "file extension not supported for thumbnail")
	ErrFileRead                  = NewAPIError(400, "file read error")
)
