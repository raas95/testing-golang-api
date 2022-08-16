package test

import (
	"github.com/gin-gonic/gin"
	// "log"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type DataNich struct {
	Ttl  string `json:"ttl" binding:"required"`
	Awek string `json:"awek" binding:"required"`
}

func PostHandler(c *gin.Context) {

	var dataNich DataNich
	err := c.ShouldBindJSON(&dataNich)
	if err != nil {
		// log.Fatal(err)
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"name": "test",
		"ttl":  dataNich.Ttl,
		"bk":   dataNich,
	})
}

func RootHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"name": "test",
		"ids":  id,
	})
}
