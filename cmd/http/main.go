package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atcheri/hexarch-go/internal/infrastructure/http-server/gin"
)

func init() {
	fmt.Println("===== BEGIN init function =====")
	fmt.Println("We could do some flag parsing here for example: -db=mysql, or -ports=8080 etc...")
	fmt.Println("===== END init function =====")
}

func main() {
	app := gin.NewGinApp()
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           app,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// Initializing the http-server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down http-server...")

	// The context is used to inform the http-server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
