package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var dbURL = "https://quotes-app-5a889-default-rtdb.asia-southeast1.firebasedatabase.app/"

// GET
func getQuotes(c *gin.Context) {
	resp, err := http.Get(dbURL + "quotes.json")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]Quote
	json.Unmarshal(body, &data)

	var quotes []Quote
	for _, q := range data {
		quotes = append(quotes, q)
	}

	c.JSON(200, quotes)
}

// POST
func addQuote(c *gin.Context) {
	var quote Quote

	if err := c.BindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonData, _ := json.Marshal(quote)

	_, err := http.Post(
		dbURL+"quotes.json",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Quote added"})
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/quotes", getQuotes)
	r.POST("/quotes", addQuote)

	r.Run(":8080")
}

// CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}