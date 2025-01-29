package migration

import (
	"os"
	"strings"

	"github.com/donbarrigon/new-project/pkg/formatter"
)

// Foreign agrega una clave foranea a la tabla
// func (t *Table) Foreign(column string, options ...string) {
// 	// si esta vacio se crea uno nuevo
// 	if t.ForeignKeys == nil {
// 		t.ForeignKeys = make([]*ForeignKey, 0)
// 		// t.ForeignKeys = append(t.ForeignKeys, Foreign(column, options...))
// 		// return
// 	}

// 	// la formateo a snake case antes de la validacion para que este en el mismo formato que la ya procesada
// 	column = formatter.ToSnakeCase(column)

// 	// verifico que exista uno de ser asi lo sobreescribo
// 	for i, fk := range t.ForeignKeys {
// 		if fk.Column == column {
// 			t.ForeignKeys[i] = t.createForeign(column, options...)
// 			return
// 		}
// 	}

// 	// si no existe solo lo agrega y ya esta
// 	t.ForeignKeys = append(t.ForeignKeys, t.createForeign(column, options...))
// }

// Foreign agrega a la tabla una clave foranea \n
/*
	// esta es la forma completa
	Foreign(column string, references, on string, onDelete string, onUpdate string)
	// asi la opcion se predetermina por la posicion
	Foreign("user_id", "id", "users", "cascade", "cascade")
	// otras formas de usarla
	// agrega o crea las opcions que nesesites o te sean convenientes
	Foreign("columns" , ["references:columns", "on:table", "ondelete:cascade", "onupdate:cascade"])
	// agrega o crea las opcions que nesesites o te sean convenientes
	Foreign("columns" , ["references:columns", "on:table"])
	// si sigues las convenciones de nombres la tabla (migracion_column) las referencias se toman del valor de column
	// puedes hacerlo asi
	Foreign("column") // las referencias se toman de column
	Foreign("column", "onDelete:cascade") // las referencias se toman column y puede combinarlo con las constraint
	// las foraneas con multiples columnas ahi si debe especificar las references y la tabla
	Foreign("columns", "references:columns", "on:table", ["ondelete:cascade", "onupdate:cascade"])
	Foreign("user_id, department_id", "references:id, department_id", "on:users", "ondelete:cascade", "onupdate:cascade")
*/
func Foreign(column string, options ...string) ForeignKey {

	// si ingresa todos los datos se crea la estructura
	if len(options) == 4 {
		return ForeignKey{
			Column:    formatter.ToSnakeCase(column),
			Reference: formatter.ToSnakeCase(options[0]),
			Table:     formatter.ToTableName(options[1]),
			OnDelete:  options[2],
			OnUpdate:  options[3],
		}
	}

	// inicial el ForeignKey
	fk := ForeignKey{
		Column: formatter.ToSnakeCase(column),
	}

	// Procesar las opciones
	for _, option := range options {
		if strings.HasPrefix(option, "references:") {
			fk.Reference = formatter.ToSnakeCase(strings.TrimPrefix(option, "references:"))
		} else if strings.HasPrefix(option, "on:") {
			fk.Table = formatter.ToTableName(strings.TrimPrefix(option, "on:"))
		} else if strings.HasPrefix(option, "ondelete:") {
			fk.OnDelete = strings.TrimPrefix(option, "ondelete:")
		} else if strings.HasPrefix(option, "onupdate:") {
			fk.OnUpdate = strings.TrimPrefix(option, "onupdate:")
		}
	}

	// Si solo se pasa la columna y no hay tabla o referencias explícitas,
	// se siguen las convenciones de nombres: tabla = columna en plural, referencias = id
	if fk.Table == "" || fk.Reference == "" {
		col := strings.Split(fk.Column, "_")
		var tableName string
		var columnName string
		if len(col) > 1 {
			tableName = strings.Join(col[:len(col)-1], "_")
			columnName = col[len(col)-1]
		} else {
			tableName = formatter.ToTableName(fk.Column)
			columnName = "id"
		}
		if fk.Table == "" {
			fk.Table = formatter.ToTableName(tableName)
		}
		if fk.Reference == "" {
			fk.Reference = formatter.ToSnakeCase(columnName)
		}
	}
	// se agrega al slice
	return fk
}

func NewTable(name string, columns ...*Column) *Table {
	table := &Table{
		Name:        formatter.ToTableName(name),
		Columns:     make([]Column, 0, len(columns)),
		PrimaryKeys: make([]string, 0),
		Constraints: make(map[string]string),
		Charset:     os.Getenv("DB_CHARSET"),
		Collation:   os.Getenv("DB_COLLATION"),
	}

	// Agregar columnas y detectar claves primarias
	for _, col := range columns {
		if col != nil {
			table.Columns = append(table.Columns, *col)
			if col.PrimaryKey {
				table.PrimaryKeys = append(table.PrimaryKeys, col.Name)
			}
		}
	}

	return table
}
