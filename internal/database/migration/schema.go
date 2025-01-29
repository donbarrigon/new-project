package migration

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/donbarrigon/new-project/pkg/formatter"
)

// NewSchema crea una nueva instancia de Schema
func NewSchema(name string) *Schema {
	return &Schema{
		Name:      name,
		Tables:    make([]Table, 0),
		Charset:   os.Getenv("DB_CHARSET"),
		Collation: os.Getenv("DB_COLLATION"),
	}
}

// CreateTable agrega una tabla nueva al schema
func (s *Schema) AddTable(table *Table) error {
	if s.HasTable(table.Name) {
		return fmt.Errorf("la tabla %s ya existe", table.Name)
	}
	s.Tables = append(s.Tables, *table)
	return nil
}

func (s *Schema) AddTables(tables ...*Table) error {
	for _, table := range tables {
		if err := s.AddTable(table); err != nil {
			return err
		}
	}
	return nil
}

// AlterTable modifica reemplaza una tabla
func (s *Schema) ReplaceTable(table *Table) error {
	for i, t := range s.Tables {
		if t.Name == table.Name {
			s.Tables[i] = *table
			return nil
		}
	}
	return fmt.Errorf("la tabla %s no existe", table.Name)
}

// HasTable verifica si existe una tabla
func (s *Schema) HasTable(name string) bool {
	name = formatter.ToTableName(name)
	for _, t := range s.Tables {
		if t.Name == name {
			return true
		}
	}
	return false
}

// DropTable elimina una tabla del schema
func (s *Schema) DropTable(name string) {
	name = formatter.ToTableName(name)
	for i, t := range s.Tables {
		if t.Name == name {
			// Eliminar el elemento en la posición i
			s.Tables = append(s.Tables[:i], s.Tables[i+1:]...)
			break // Salir del bucle después de eliminar
		}
	}
}

// RenameTable renombra una tabla
func (s *Schema) RenameTable(oldName, newName string) error {

	oldName = formatter.ToTableName(oldName)
	newName = formatter.ToTableName(newName)

	for _, t := range s.Tables {
		if t.Name == newName {
			return fmt.Errorf("la tabla %s ya existe", newName)
		}
	}

	for _, t := range s.Tables {
		if t.Name == oldName {
			t.Name = formatter.ToTableName(newName)
			return nil
		}
	}
	return fmt.Errorf("tabla %s no existe", oldName)
}

// ApplyDefaults aplica la configuración por defecto del schema a las tablas que no tienen configuración específica
func (s *Schema) ApplyDefaults() {
	for _, table := range s.Tables {
		if table.Charset == "" {
			table.Charset = s.Charset
		}
		if table.Collation == "" {
			table.Collation = s.Collation
		}
	}
}

// GetTableNames retorna los nombres de todas las tablas en el schema
func (s *Schema) GetTableNames() []string {
	names := make([]string, 0, len(s.Tables))
	for _, t := range s.Tables {
		names = append(names, t.Name)
	}
	sort.Strings(names)
	return names
}

// Validate verifica la integridad del schema
func (s *Schema) Validate() error {
	errors := make([]string, 0)

	// Verifica nombres únicos de tablas
	for _, table := range s.Tables {
		for _, t := range s.Tables {
			if table.Name == t.Name {
				errors = append(errors, fmt.Sprintf("nombre de tabla duplicado: %s", table.Name))
			}
		}
	}

	// Verifica referencias de claves foráneas
	// for tableName, table := range s.Tables {
	// 	for _, fk := range table.ForeignKeys {
	// 		// Verifica que la tabla referenciada exista
	// 		refTable, exists := s.Tables[fk.Table]
	// 		if !exists {
	// 			errors = append(errors, fmt.Sprintf("tabla %s: referencia a tabla inexistente %s", tableName, fk.Table))
	// 			continue
	// 		}

	// 		// Verifica que la columna referenciada exista
	// 		columnExists := false
	// 		for _, col := range refTable.Columns {
	// 			if col.Name == fk.Reference {
	// 				columnExists = true
	// 				break
	// 			}
	// 		}
	// 		if !columnExists {
	// 			errors = append(errors, fmt.Sprintf("tabla %s: referencia a columna inexistente %s en tabla %s",
	// 				tableName, fk.Reference, fk.Table))
	// 		}
	// 	}
	// }

	if len(errors) > 0 {
		return fmt.Errorf("errores de validación:\n%s", strings.Join(errors, "\n"))
	}

	return nil
}
