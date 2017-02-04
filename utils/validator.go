package utils

import (
	"reflect"
	"gopkg.in/gin-gonic/gin.v1"
)

// Validating and Binding form data
func Validator(v interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := reflect.New(reflect.TypeOf(v)).Interface()
		err := c.Bind(a)
		if err != nil {
			respondWithError(401, err.Error(), c)
			return
		}
		c.Set("form_data", a)
		c.Next()
	}
}



func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

