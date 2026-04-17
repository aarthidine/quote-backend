package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// 📌 Quote model
type Quote struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// 📌 in-memory storage
var quotes = []Quote{
	{ID: "1", Text: "Stay hungry, stay foolish", Author: "Steve Jobs"},
}

// 🚀 MAIN FUNCTION
func main() {
	r := gin.Default()

	// ✅ CORS FIX (VERY IMPORTANT FOR VERCEL + MOBILE)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://quote-frontend-zukp.vercel.app",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},
	}))

	// 📌 Routes
	r.GET("/quotes", getQuotes)
	r.POST("/quotes", addQuote)
	r.DELETE("/quotes/:id", deleteQuote)

	// 🚀 start server
	r.Run(":8080")
}