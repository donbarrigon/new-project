package main

import (
	"log"
	"os"

	"github.com/erespereza/new-project/config"
	"github.com/erespereza/new-project/internal/app"
	"github.com/erespereza/new-project/internal/orm"
)

func main() {

	// Carga las variables del archivo .env
	config.Load()

	// Conecta con la base de datos
	orm.Connect()

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
