package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DB es la instancia global de la base de datos
var db *sql.DB

// Connect establece una conexión con la base de datos
func Connect() {
	// Obtiene las variables de la base de datos desde el entorno (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_CHARSET, DB_COLLATION)
	// Recomendación: Utilizar variables de entorno en lugar de hardcodear los datos de conexión a la base de datos para proteger la información sensible.
	// Consulte https://12factor.net/config para más información sobre cómo se deben manejar las variables de entorno en aplicaciones Go.

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "mysql"
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbCollation := os.Getenv("DB_COLLATION")

	// Construye la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbCollation)

	// Abre la conexión con la base de datos
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión con la base de datos: %v", err)
	}

	// Verifica que la conexión sea válida
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al verificar la conexión con la base de datos: %v", err)
	}

	log.Println("Conexión con la base de datos establecida.")
}
