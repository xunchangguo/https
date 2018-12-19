package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "HEAD, PATCH, POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("options")
			return
		}
		c.Next()
	}
}

func main() {
	httpAddr := flag.String("listen", ":9090", "address:port to listen on")
	dir := flag.String("dir", "", "static file dir")
	certFile := flag.String("certFile", "server.pem", "https certFile")
	keyFile := flag.String("keyFile", "server.key", "https keyFile")
	flag.Parse()

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.StaticFS("/", http.Dir(*dir))

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(*httpAddr, *certFile, *keyFile)
}
