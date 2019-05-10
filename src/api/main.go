package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml-gorutines/src/api/controllers/myml"
	"github.com/mercadolibre/myml-gorutines/src/api/controllers/ping"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", ping.Ping)
	router.GET("/myml/:userID", myml.GetUser)
	router.Run(port)
}
