package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erespereza/new-project/internal/routes"
)

// Response estructura para respuestas HTTP estandarizadas
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// handleHealthCheck maneja el endpoint de health check
func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handleMethodNotAllowed(w, r)
		return
	}

	response := Response{
		Status:  "success",
		Message: "Service is healthy",
		Data: map[string]string{
			"version": "1.0.0",
			"status":  "up",
		},
	}

	sendJSONResponse(w, http.StatusOK, response)
}

// handleMethodNotAllowed maneja métodos HTTP no permitidos
func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "error",
		Message: fmt.Sprintf("Method %s not allowed", r.Method),
	}
	sendJSONResponse(w, http.StatusMethodNotAllowed, response)
}

// sendJSONResponse envía una respuesta JSON estandarizada
func sendJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

// loggingMiddleware implementa logging para cada request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next(w, r)

		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}

// setupRoutes configura las rutas del servidor
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", loggingMiddleware(handleHealthCheck))

	// API v1 endpoints
	// mux.HandleFunc("/api/v1/users", loggingMiddleware(handleGetUsers))
	// mux.HandleFunc("/api/v1/users/create", loggingMiddleware(handleCreateUser))
	// mux.HandleFunc("/api/v1/products", loggingMiddleware(handleGetProducts))
	return mux
}

// StartServer inicia el servidor HTTP
func StartServer(port string) *http.Server {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      routes.NewRouter(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Inicia el servidor en una goroutine separada
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	return server
}

// gracefulShutdown maneja el apagado graceful del servidor
func GracefulShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Espera por la señal de terminación
	<-sigChan
	log.Println("Shutting down server...")

	// Crea un contexto con timeout para el shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}
