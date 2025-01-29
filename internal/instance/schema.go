package instance

import (
	"github.com/donbarrigon/new-project/internal/database/migration"
)

// variable donde se va a almacenar la estructura de la base de datos
// para que pueda ser accedida desde cualquier modelo en cualquier parte
// sin ir a buscar a la base de datos ni procesar nada que este ahi ala mano
var schema migration.Schema

func GetTable(name string) *migration.Table {
	
}

func GetSchema() *migration.Schema {
	return &schema
}

func NewSchema(db *migration.Schema) {
	schema = *db
}
