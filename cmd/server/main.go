package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kazuyainoue0124/go-rest-api/config"
	"github.com/kazuyainoue0124/go-rest-api/infrastructure/db"
	"github.com/kazuyainoue0124/go-rest-api/infrastructure/repository"
	"github.com/kazuyainoue0124/go-rest-api/presentation/handlers"
	"github.com/kazuyainoue0124/go-rest-api/presentation/router"
	"github.com/kazuyainoue0124/go-rest-api/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	database, err := db.NewMySQLDB(cfg)
	if err != nil {
		log.Fatal("failed to connect db: ", err)
	}

	repo := repository.NewMySQLTaskRepository(database)
	u := usecase.NewTaskUsecase(repo)
	h := handlers.NewTaskHandler(u)

	mux := router.NewRouter(h)
	portStr := ":" + strconv.Itoa(cfg.App.Port)

	srv := &http.Server{
		Addr:    portStr,
		Handler: mux,
	}

	go func() {
		log.Println("server started on port", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	if err := srv.Close(); err != nil {
		log.Fatal("failed to shutdown server gracefully", err)
	}
	log.Println("server exited")
}
