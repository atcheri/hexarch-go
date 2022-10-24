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

	"github.com/atcheri/hexarch-go/internal/core/adapters/right/repositories/inMemory"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
	server "github.com/atcheri/hexarch-go/internal/infrastructure/http-server/gin"
)

var (
	port = "8080"
)

func init() {
	fmt.Println("===== BEGIN init function =====")
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Println("===== END init function =====")
}

func main() {
	// Initializing the DB
	db := databases.NewInMemoryDB()

	// Initializing repositories
	translationsRepo := adapters.NewInMemoryTranslations(db)

	app := server.NewGinApp(server.NewAppControllers(server.AppControllersDependencies{TranslationsRepo: translationsRepo}))

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           app,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// Initializing the http-server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		fmt.Printf("Server running and listening on port: %s\n", port)
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
