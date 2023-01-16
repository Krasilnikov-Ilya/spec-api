package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", getSpec)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getSpec(c *gin.Context) {
	f, err := os.ReadFile("openapi.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)

	if err != nil {
		log.Fatal(err)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.JSON(http.StatusOK, data)
}
