package instance

import "github.com/donbarrigon/new-project/internal/database/migration"

// variable donde se va a almacenar la estructura de la base de datos para que pueda ser accedida desde cualquier modelo en cualquier parte
var schema map[string]migration.Table

func GetMigration(nameMigration string) migration.Table {
	return schema[nameMigration]
}

func SetMigration(nameMigration string, migration migration.Table) {
	schema[nameMigration] = migration
}

func NewMigration(migrations ...map[string]migration.Table) {
	schema = make(map[string]migration.Table)
}
