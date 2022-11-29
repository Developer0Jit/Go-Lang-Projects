package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PingHandler(c *gin.Context) {
	c.AbortWithStatusJSON(200, gin.H{
		"message": "Hi, this is from inside the server ping route",
	})
}

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func JsonLoadHandler(c *gin.Context) {
	data := []Employee{}
	file, err := os.Open("MOCK_DATA.json")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	byt, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if json.Unmarshal(byt, &data) != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Set("jsonemployees", data)
	log.Debugf("We have about %d employees loaded", len(data))
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(false)
	log.SetLevel(log.DebugLevel)
	log.Info("Starting the go gin application")
	defer log.Warn("Shutting down the gin application server")

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", PingHandler)
	employees := r.Group("/employees", JsonLoadHandler)
	employees.GET("/", func(ctx *gin.Context) {
		val, _ := ctx.Get("jsonemployees")
		data := val.([]Employee)
		ctx.AbortWithStatusJSON(200, gin.H{
			"total": len(data),
		})
	})
	log.Fatal(r.Run(":9090"))
}
