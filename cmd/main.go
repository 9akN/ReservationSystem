package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourname/reservation-system/config"
	"github.com/yourname/reservation-system/db"
	"github.com/yourname/reservation-system/internal/handlers"
	"github.com/yourname/reservation-system/internal/repository"
	"github.com/yourname/reservation-system/internal/service"
)

func main() {
	cfg := config.Load()

	// Run database migrations
	if err := db.RunMigrations(cfg.DSN()); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("migrations applied successfully")

	// Connect to database
	pool, err := db.NewPool(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("connected to database")

	// Wire up layers
	reservationRepo := repository.NewReservationRepository(pool)
	reservationSvc := service.NewReservationService(reservationRepo)
	reservationHandler := handlers.NewReservationHandler(reservationSvc)

	// Register routes
	mux := http.NewServeMux()
	reservationHandler.RegisterRoutes(mux)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
