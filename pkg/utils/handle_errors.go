package utils

import (
	"clean-architecture/pkg/api_errors"
	"clean-architecture/pkg/framework"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func HandleValidationError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func HandleErrorWithStatus(logger framework.Logger, c *gin.Context, statusCode int, err error) {
	logger.Error(err)
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

// list static errors to filter
var exceptStaticError = []error{}

// list dyanmic errors to filter
var exceptDynamicError = []error{}

// list SQL errors to filter
var exceptSQLError = []uint16{
	1062, // duplicate entry
}

var sqlError *mysql.MySQLError

func HandleError(logger framework.Logger, c *gin.Context, err error) {
	logger.Error(err)

	msgForUnhandledError := "An error occurred while processing your request. Please try again later."

	if apiErr, ok := err.(*api_errors.APIError); ok {
		c.JSON(apiErr.StatusCode, gin.H{
			"error": apiErr.Message,
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": gorm.ErrRecordNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": msgForUnhandledError,
	})

	for _, e := range exceptStaticError {
		if errors.Is(err, e) {
			return
		}
	}

	for _, e := range exceptDynamicError {
		if reflect.TypeOf(e) == reflect.TypeOf(err) {
			return
		}
	}

	if errors.As(err, &sqlError) {
		for _, code := range exceptSQLError {
			if code == sqlError.Number {
				return
			}
		}
	}

	CurrentSentryService.CaptureException(err)
}
