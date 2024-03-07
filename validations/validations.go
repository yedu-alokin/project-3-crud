package validations

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func ValidateRequestBody(c *gin.Context) {
    var req CreateStudentValidator
    fmt.Println("--------------------------1----------------------------------")
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        c.Abort()
        return
    }
    fmt.Println("--------------------------2----------------------------------")


    // Use the validator package to perform additional validation
    validate := validator.New()
    if err := validate.Struct(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        c.Abort()
        return
    }
    fmt.Println("--------------------------3----------------------------------")


    // If validation passes, continue to the next handler
    // c.Set("myRequest", req)
    c.Next()
}