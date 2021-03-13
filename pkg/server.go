package pkg

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"regexp"
)

func SetupRouter() *gin.Engine {
	// Logging setup
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	// Create a gin.Engine
	r := gin.Default()

	r.GET("/:word", HelloServer)

	// Not allowed methods
	r.POST("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})
	r.PUT("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})
	r.DELETE("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})
	r.PATCH("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})
	r.HEAD("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})
	r.OPTIONS("/:word", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed"})
	})

	// Change default 404 to be JSON
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	return r
}

func HelloServer(c *gin.Context) {
	// Sample logging
	log.Info("Info log is just gibberish :)...")
	log.Warn("I warned you ...")
	log.Error("... That it would be an error.")

	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	if re.MatchString(c.Param("word")) == false {
		c.AbortWithStatusJSON(400, gin.H{"message": "only alphanumeric characters are allowed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": c.Param("word")})
}
