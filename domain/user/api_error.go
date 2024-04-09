package user

import (
	"clean-architecture/pkg/api_errors"
	"net/http"
)

var (
	ErrInvalidUserID = api_errors.NewAPIError(http.StatusBadRequest, "Invalid user ID")
)
