package migration

import (
	"strings"

	"github.com/donbarrigon/new-project/pkg/formatter"
)

// Foreign agrega a la tabla una clave foranea
// Foreign(column string, references, on string, onDelete string, onUpdate string) esta es la forma completa
// Foreign("user_id", "id", "users", "cascade", "cascade") // asi la opcion se determina por la posicion
// otras formas de usarla
// Foreign("columns" , ["references:columns", "on:table", "ondelete:cascade", "onupdate:cascade"]) agrega o crea las opcions que nesesites o te sean convenientes
// Foreign("columns" , ["references:columns", "on:table"]) agrega o crea las opcions que nesesites o te sean convenientes
// si sigues las convenciones de nombres la tabla (migracion_column) las referencias se toman del valor de column
// puedes hacerlo asi
// Foreign("column") // las referencias se toman de column
// Foreign("column", "onDelete:cascade") // las referencias se toman column y puede combinarlo con las constraint
// las foraneas con multiples columnas ahi si debe especificar las references y la tabla
// Foreign("columns", "references:columns", "on:table", ["ondelete:cascade", "onupdate:cascade"])
// Foreign("user_id, department_id", "references:id, department_id", "on:users", "ondelete:cascade", "onupdate:cascade")
func (t *Table) Foreign(column string, options ...string) {
	// si ingresa todos los datos se crea la estructura
	if len(options) == 4 {
		t.ForeignKeys = append(t.ForeignKeys, ForeignKey{
			Column:     formatter.ToSnakeCase(column),
			References: formatter.ToSnakeCase(options[0]),
			On:         formatter.ToTableName(options[1]),
			OnDelete:   options[2],
			OnUpdate:   options[3],
		})
	}

	// inicial el ForeignKey
	fk := ForeignKey{
		Column: formatter.ToSnakeCase(column),
	}

	// Procesar las opciones
	for _, option := range options {
		if strings.HasPrefix(option, "references:") {
			fk.References = formatter.ToSnakeCase(strings.TrimPrefix(option, "references:"))
		} else if strings.HasPrefix(option, "on:") {
			fk.On = formatter.ToTableName(strings.TrimPrefix(option, "on:"))
		} else if strings.HasPrefix(option, "ondelete:") {
			fk.OnDelete = strings.TrimPrefix(option, "ondelete:")
		} else if strings.HasPrefix(option, "onupdate:") {
			fk.OnUpdate = strings.TrimPrefix(option, "onupdate:")
		}
	}

	// Si solo se pasa la columna y no hay tabla o referencias explícitas,
	// se siguen las convenciones de nombres: tabla = columna en plural, referencias = id
	if fk.On == "" || fk.References == "" {
		col := strings.Split(column, "_")
		var tableName string
		var columnName string
		if len(col) > 1 {
			tableName = strings.Join(col[:len(col)-1], "_")
			columnName = col[len(col)-1]
		} else {
			tableName = formatter.ToTableName(column)
			columnName = "id"
		}
		if fk.On == "" {
			fk.On = formatter.ToTableName(tableName)
		}
		if fk.References == "" {
			fk.References = formatter.ToSnakeCase(columnName)
		}
	}

	// se agrega al slice
	t.ForeignKeys = append(t.ForeignKeys, fk)
}
