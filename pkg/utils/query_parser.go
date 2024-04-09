package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryOption struct {
	Name       string
	Type       reflect.Type
	IsRequired bool
}

/*
Parses the query parameters from the given list of QueryOption and the gin.Context.
It returns a map[string]interface{} containing the parsed query parameters and an error if any.
The map keys are the names of the query parameters and the values are the parsed values.
If a query parameter is missing and marked as required, it returns an error.
The query parameter values are transformed based on their specified types in the QueryOption.
Supported types are string, int, uint, and bool.
If the type is not supported, the value is stored as a string

Usage example:

	queryData, err := utils.ParseQueries([]utils.QueryOption{{Name: "search", Type: reflect.TypeOf(false), IsRequired: true}}, ctx)
	if err != nil {
		utils.HandleError(c.logger, ctx, err)
		return
	}
	searchKey, _ := queryData["search"].(bool)
*/
func ParseQueries(queries []QueryOption, ctx *gin.Context) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})

	for _, query := range queries {
		val := ctx.Query(query.Name)

		if val == "" {
			if query.IsRequired {
				return nil, fmt.Errorf("ParseQueries(): %s is missing", query.Name)
			} else {
				continue
			}
		}

		var transformedVal interface{}
		var err error

		switch query.Type.Kind() {

		case reflect.String:
			transformedVal = val

		case reflect.Int:
			transformedVal, err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}

		case reflect.Uint:
			transformedVal, err = StringToUInt(val)
			if err != nil {
				return nil, err
			}

		case reflect.Bool:
			if val == "true" {
				transformedVal = true
			} else {
				transformedVal = false
			}

		default:
			transformedVal = val
		}

		dataMap[query.Name] = transformedVal
	}

	return dataMap, nil
}
