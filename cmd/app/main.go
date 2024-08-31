package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routes "github.com/proyectum/ms-user-profile/internal/adapters/in/http"
	"github.com/proyectum/ms-user-profile/internal/boot"
	"log"
	"net/http"
	"time"
)

func main() {
	boot.LoadConfig()
	boot.ExecuteMigrations()
	r := gin.Default()
	r.GET("/ping", ping)
	routes.RegisterRoutes(r)
	s := createServer(r)
	log.Fatal(s.ListenAndServe())
}

func createServer(r *gin.Engine) *http.Server {
	serverProps := boot.CONFIG.Server
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", serverProps.Port),
		Handler:        r,
		ReadTimeout:    serverProps.ReadTimeout * time.Second,
		WriteTimeout:   serverProps.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
