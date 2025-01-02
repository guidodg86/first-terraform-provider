package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Element struct {
	ID          string `json:"id"`
	Number      int    `json:"number"`
	String_data string `json:"string_data"`
}

type NewElement struct {
	Number      int    `json:"number"`
	String_data string `json:"string_data"`
}

var Elements []Element

// returns all data
func getAllElements(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Elements)
}

// Create new element on memory
func createElement(c *gin.Context) {
	var ReceivedData NewElement

	//To begin three error condition check
	if err := c.ShouldBindJSON(&ReceivedData); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if ReceivedData.Number == 0 {
		c.Error(errors.New("example-api: number 0 is not allowed - check fields name on json sent"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if ReceivedData.String_data == "" {
		c.Error(errors.New("example-api: empty string are not allowed - check fields name on json sent"))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//We can save new element on memory and return ok
	var current_time = time.Now()
	var new_hash = md5.Sum([]byte(current_time.String()))
	AddedElement := Element{
		ID:          hex.EncodeToString(new_hash[:]),
		Number:      ReceivedData.Number,
		String_data: ReceivedData.String_data}
	Elements = append(Elements, AddedElement)
	c.IndentedJSON(http.StatusOK, ReceivedData)
}

func main() {

	//Creating first element for testing reasons
	element_example := Element{
		ID:          "c930e4d1ec2b1f0035f85038b055d644",
		Number:      25,
		String_data: "Velez Campeon 2024"}
	Elements = append(Elements, element_example)

	// Initializing GIN router
	router := gin.Default()
	router.GET("/all_data", getAllElements)
	router.POST("/create", createElement)

	//Starting the server
	router.Run("localhost:8080")
}
