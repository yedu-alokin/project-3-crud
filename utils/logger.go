package utils

import (
	"crud/models/objects"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Looger(c *gin.Context, message interface{}) error {

	logData := objects.Logs{
		URL:      c.Request.RequestURI,
		Method:   c.Request.Method,
		Response: "test",
		Message:  message,
	}

	file, err := os.OpenFile("crud.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("an error happened while checking for log file")
		return err
	}
	defer file.Close()

	// Create a new logger that writes to the file
	logger := log.New(file, "", log.LstdFlags)

	// Write the message to the log file
	logger.Println(logData)

	return nil
}
