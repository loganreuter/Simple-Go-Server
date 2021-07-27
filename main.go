package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port string = ":3000"

func main() {
	//Makes a router
	r := gin.Default()

	//Path handling
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	r.GET("/HelloWorld", func(c *gin.Context) {
		c.File("./index.html")
	})

	//Creates server object
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//Starts server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started: http://localhost" + port)

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
