package tables

import (
	"os"

	. "github.com/donbarrigon/new-project/internal/database/migration"
	"github.com/donbarrigon/new-project/internal/instance"
)

// NewMigration crea la migracion y ademas 
func NewMigration() {
	// toma el nombre de la tabla del .env
	schema := NewSchema(os.Getenv("DB_NAME"))

	schema.AddTables(
		create_user_table(),
		// aqui agrege las demas funciones de creacion de tablas
	)

	// esta linea es la ultima si va a modificar algo del schema agalo antes de esta linea
	// registra el schema para que los modelos accedan a el cuando lo nesesiten
	instance.NewSchema(schema)
}
