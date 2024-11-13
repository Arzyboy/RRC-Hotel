package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Sajikan file statis untuk CSS
	r.Static("/static", "./static")

	// Setel template HTML
	r.LoadHTMLGlob("templates/*.html")

	// Routing untuk halaman Utama
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout.html", gin.H{
			"content": "rooms.html",
		})
	})

	r.Run(":8080")
}
