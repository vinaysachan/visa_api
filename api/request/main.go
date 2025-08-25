package request

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaysachan/visa_api/base/utils"
)

// genericValidate handles common validation logic for all request structs
func GenericValidate(c *gin.Context, request interface{}) bool {
	// Bind JSON Validation
	err := c.ShouldBindJSON(request)
	if errors.Is(err, io.EOF) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty body, please send some data"})
		return false
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return false
	}

	// Request Struct Validation
	serr := utils.Validate.Struct(request)
	if serr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": utils.NewValidationError(serr).Errors})
		return false
	}

	return true
}
