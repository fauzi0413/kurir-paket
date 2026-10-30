package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kurir-paket/backend/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverHandler := api.NewServerHandler()
	mux := http.NewServeMux()

	// REST API Endpoints
	mux.HandleFunc("/health", serverHandler.HealthHandler)
	mux.HandleFunc("/api/v1/vehicles", serverHandler.GetVehiclesHandler)
	mux.HandleFunc("/api/v1/dispatch/next", serverHandler.DispatchNextHandler)
	mux.HandleFunc("/api/v1/telemetry/verify", serverHandler.TelemetryVerifyHandler)
	mux.HandleFunc("/api/v1/player/save", serverHandler.PlayerSaveHandler)
	mux.HandleFunc("/api/v1/player/load", serverHandler.PlayerLoadHandler)
	mux.HandleFunc("/api/v1/leaderboard", serverHandler.LeaderboardHandler)

	// Opsi Static File Serving (Monolithic Single-Container Deployment)
	serveStatic := os.Getenv("SERVE_STATIC")
	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "../frontend"
	}

	if serveStatic == "true" {
		if _, err := os.Stat(staticDir); err == nil {
			log.Printf(" [Static] Serving frontend static files from: %s", staticDir)
			fileServer := http.FileServer(http.Dir(staticDir))
			mux.Handle("/", fileServer)
		} else {
			log.Printf(" [Static Warning] Static directory not found: %s", staticDir)
		}
	}

	addr := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel untuk graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("=====================================================")
		log.Printf(" Warehouse Express 3D - Go Game Server Running")
		log.Printf(" Listening on http://localhost%s", addr)
		log.Printf(" Health check: http://localhost%s/health", addr)
		if serveStatic == "true" {
			log.Printf(" Static UI:   http://localhost%s/", addr)
		}
		log.Printf("=====================================================")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Tunggu sinyal terminate / interrupt
	<-stop
	log.Println("Shutting down game server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Game server stopped cleanly.")
}
