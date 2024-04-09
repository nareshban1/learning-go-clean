package utils

import (
	"clean-architecture/pkg/api_errors"
	"clean-architecture/pkg/framework"
	"errors"
	"net/http"
	"reflect"

	"github.com/getsentry/sentry-go"
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

	// will not captured by sentry if its an explicit APIError
	if apiErr, ok := err.(*api_errors.APIError); ok {
		c.JSON(apiErr.StatusCode, gin.H{
			"error": apiErr.Message,
		})
		return
	}

	// will not captured by sentry if its an record not found error from gorm
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
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

	sentry.CaptureException(err)
}
