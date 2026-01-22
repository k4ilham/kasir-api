package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kasir-api/handlers"
	"kasir-api/response"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Running on port :8085"))
}

func health2Handler(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/produk/", handlers.ProdukItemHandler)
	mux.HandleFunc("/api/produk", handlers.ProdukCollectionHandler)
	mux.HandleFunc("/api/categories/", handlers.CategoryItemHandler)
	mux.HandleFunc("/api/categories", handlers.CategoryCollectionHandler)
	mux.HandleFunc("/health", health2Handler)

	server := &http.Server{
		Addr:    ":8085",
		Handler: mux,
	}

	go func() {
		log.Println("Server running di localhost:8085")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("gagal running server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
