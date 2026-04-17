import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
)

func main() {
    r := gin.Default()

    // ✅ CORS FIX (IMPORTANT)
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "https://quote-frontend-zukp.vercel.app",
            "http://localhost:3000",
        },
        AllowMethods: []string{
            "GET",
            "POST",
            "DELETE",
            "PUT",
            "OPTIONS",
        },
        AllowHeaders: []string{
            "Origin",
            "Content-Type",
        },
        MaxAge: 12 * time.Hour,
    }))

    // Routes
    r.GET("/quotes", getQuotes)
    r.POST("/quotes", addQuote)
    r.DELETE("/quotes/:id", deleteQuote)

    r.Run(":8080")
}