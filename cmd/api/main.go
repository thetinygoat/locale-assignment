package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/thetinygoat/localeai/pkg/queue"
	_rideHandler "github.com/thetinygoat/localeai/pkg/ride/handler/http"
	_rideRepo "github.com/thetinygoat/localeai/pkg/ride/repository/postgres"
	_rideService "github.com/thetinygoat/localeai/pkg/ride/service"
	"github.com/thetinygoat/localeai/pkg/worker"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// setup database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	q := queue.NewQueue("rideDump")
	consumer, err := q.AddConsumer("rideConsumer")
	if err != nil {
		log.Fatal(err)
	}
	rideRepo := _rideRepo.New(conn)
	rideService := _rideService.New(rideRepo)
	// register routes
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/ride", _rideHandler.Routes(q))
	})
	go worker.StartWorkers(10, consumer, rideService)
	// start the server
	log.Println("starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
