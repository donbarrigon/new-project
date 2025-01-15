package main

import (
	"log"
	"os"

	"github.com/erespereza/clan-de-raid/internal/app"
)

func main() {
	// Configura el logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicia el servidor
	server := app.StartServer(port)

	// Maneja el shutdown graceful
	app.GracefulShutdown(server)
}
