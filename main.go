package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
	http2 "stvcv2/http"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {

	router := gin.New()

	//server := network.GetServer()
	//network.Start()
	//go server.Serve()
	//defer server.Close()

	fmt.Println("Started:")

	router.Use(GinMiddleware("http://127.0.0.1:5500"))
	router.POST("/run-command", gin.WrapF(http2.HandleRunCommand))

	//router.GET("/socket.io/*any", gin.WrapH(server))
	//router.POST("/socket.io/*any", gin.WrapH(server))

	if err := router.Run(":8081"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
