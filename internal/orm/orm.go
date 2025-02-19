package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// db es la instancia global de la base de datos
var db *sql.DB

// MongoDBClient representa la conexión a la base de datos
var dbc *mongo.Client
var databaseName string
var dbDriver string

// Connect establece una conexión con la base de datos
func Connect() {

	dbDriver = os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "mongodb"
	}

	if dbDriver == "mongodb" {
		connectMongoDB()
		return
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

// ConnectDB inicializa la conexión con MongoDB
func connectMongoDB() *mongo.Client {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener URI y nombre de la base de datos desde .env
	mongoURI := os.Getenv("MONGO_URI")

	// Configurar opciones de conexión
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar a MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	// Verificar la conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}

	dbc = client
	databaseName = os.Getenv("DB_NAME")

	fmt.Println("Conectado a MongoDB", dbc)
	return client
}

// Find busca un registro en la base de datos basado en el ID
// id puede ser int, float o string
// columns es opcional, si no se proporciona se usan todos los campos del struct
// se valida que almenos una de las columnas se valida para la busqueda

// func (e *ExtendsModel) scanRows(m *Model, row *sql.Row, columns []string) error {

// 	// obtener el valor dinámico del modelo
// 	modelValue := reflect.ValueOf(m)
// 	// Obtener el tipo del struct
// 	structType := modelValue.Elem().Type()

// 	// Crear slice para almacenar los valores de las columnas
// 	// se hacen dos slices uno con los valores otro con las referencias por que a row.Scan hay que pasarle punteros
// 	values := make([]interface{}, len(columns))
// 	scanRefs := make([]interface{}, len(columns))
// 	for i := range values {
// 		scanRefs[i] = &values[i]
// 	}

// 	// Escanear resultados
// 	if err := row.Scan(scanRefs...); err != nil {
// 		if err == sql.ErrNoRows {
// 			return fmt.Errorf("registro no encontrado")
// 		}
// 		return fmt.Errorf("error al escanear resultados: %v", err)
// 	}

// 	// Asignar valores al struct
// 	structValue := modelValue.Elem()
// 	for i, colName := range columns {
// 		// Convertir nombre de columna a PascalCase para que coincida con el del struct
// 		colName = formatter.ToPascalCase(colName)

// 		// Buscar el campo correspondiente en el struct
// 		var field reflect.Value
// 		for j := 0; j < structType.NumField(); j++ {
// 			if structType.Field(j).Name == colName {
// 				field = structValue.Field(j)
// 				break
// 			}
// 		}

// 		if field.IsValid() && field.CanSet() {
// 			// Convertir y asignar el valor
// 			val := reflect.ValueOf(values[i])
// 			if val.Type().ConvertibleTo(field.Type()) {
// 				field.Set(val.Convert(field.Type()))
// 			}
// 			// aca hay un detalle que si no se puede convertir queda nulo ayayayay
// 		}
// 	}

// 	return nil
// }
