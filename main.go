package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 🔥 YOUR FIREBASE URL
const firebaseURL = "https://quotes-app-5a889-default-rtdb.asia-southeast1.firebasedatabase.app"

// 📌 Quote model
type Quote struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	r := gin.Default()

	// ✅ CORS (for Vercel + mobile)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://quote-frontend-zukp.vercel.app",
		},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// 📌 Routes
	r.GET("/quotes", getQuotes)
	r.POST("/quotes", addQuote)
	r.DELETE("/quotes/:id", deleteQuote)

	// 🔥 Render PORT fix
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

//
// ==============================
// 📌 GET QUOTES (FIXED WITH ID)
// ==============================
//
func getQuotes(c *gin.Context) {
	resp, err := http.Get(firebaseURL + "/quotes.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]Quote
	json.Unmarshal(body, &data)

	var quotes []Quote

	// 🔥 IMPORTANT: Add ID from Firebase key
	for key, q := range data {
		q.ID = key
		quotes = append(quotes, q)
	}

	c.JSON(200, quotes)
}

//
// ==============================
// 📌 ADD QUOTE
// ==============================
//
func addQuote(c *gin.Context) {
	var quote Quote

	if err := c.BindJSON(&quote); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	jsonData, _ := json.Marshal(quote)

	_, err := http.Post(firebaseURL+"/quotes.json", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add"})
		return
	}

	c.JSON(200, gin.H{"message": "Quote added"})
}

//
// ==============================
// 📌 DELETE QUOTE
// ==============================
//
func deleteQuote(c *gin.Context) {
	id := c.Param("id")

	req, _ := http.NewRequest("DELETE", firebaseURL+"/quotes/"+id+".json", nil)
	client := &http.Client{}
	_, err := client.Do(req)

	if err != nil {
		c.JSON(500, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Deleted"})
}