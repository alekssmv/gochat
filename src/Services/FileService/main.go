package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Serve static files.
	r.Static("/public", "./public")

  // Routes. 
  r.GET("/login", func(c *gin.Context) {
    // Get html from public folder.
    c.File("./public/login.html")
    // Return status code 200.
    c.Status(http.StatusOK)
  })

  r.GET("/register", func(c *gin.Context) {
    // Get html from public folder.
    c.File("./public/register.html")
    // Return status code 200.
    c.Status(http.StatusOK)
  })

  r.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "/login")
  })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
