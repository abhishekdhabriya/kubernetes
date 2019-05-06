package main

import (
	"gohello/routers"
	"gohello/plugins"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/afex/hystrix-go/hystrix/metric_collector"
	log "github.com/sirupsen/logrus"
	"os"
	"net/http"
)

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func HystrixHandler(command string) gin.HandlerFunc {
	return func(c *gin.Context) {
		hystrix.Do(command, func() error {
			c.Next()
			return nil
		}, func(err error) error {
			c.String(http.StatusInternalServerError, "500 Internal Server Error")
			return err
		})
	}
}


func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)


	// Hystrix configuration
	hystrix.ConfigureCommand("timeout", hystrix.CommandConfig{
		Timeout: 1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})
	

	router := gin.Default()
	router.RedirectTrailingSlash = false


	router.Use(HystrixHandler("timeout"))

	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "You are now running a blank Go application from http://leanpub.com/kube")
	})
	router.GET("/health", routers.HealthGET)

	log.Info("Starting gohello on port " + port())

	router.Run(port())
}
