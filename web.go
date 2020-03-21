package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func startWeb() {
	r := setupRouter()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go server.ListenAndServe()
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logFormatter), gin.Recovery())

	r.POST("/alert", func(c *gin.Context) {
		var alert Alert

		if err := c.ShouldBindJSON(&alert); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		queue <- alert

		c.JSON(http.StatusOK, gin.H{"status": "nominal"})
	})

	return r
}

func logFormatter(param gin.LogFormatterParams) string {
	m := map[string]interface{}{
		"type":       "access_log",
		"client_ip":  param.ClientIP,
		"time":       param.TimeStamp.Format(time.RFC3339),
		"method":     param.Method,
		"path":       param.Path,
		"protocol":   param.Request.Proto,
		"status":     param.StatusCode,
		"latency":    param.Latency,
		"user_agent": param.Request.UserAgent(),
		"error":      param.ErrorMessage,
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error while marshalling to JSON:", err)
		return "{}"
	}

	return fmt.Sprintf("%s\n", bytes)
}
