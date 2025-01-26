package orm

import (
	"fmt"
	"strings"
)

func (e *Model) Find(id any, columns ...string) error {

	// Validar tipo de ID
	switch id.(type) {
	case int, int8, int16, int32, int64, string:
		// tipos válidos no se hace nada
	default:
		return fmt.Errorf("id debe ser numérico o string")
	}

	// Determinar las columnas a consultar
	if len(columns) > 0 {
		// Usar las columnas proporcionadas
		e.selectColumns = columns
	} else {
		// Usar las columnas del modelo
		e.selectColumns = e.columns
	}

	// Si no hay columnas seleccionadas, retornar error
	if len(e.selectColumns) == 0 {
		return fmt.Errorf("no hay campos válidos para consultar")
	}

	// Construir la consulta SQL
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id = $1",
		strings.Join(e.selectColumns, ", "),
		e.tableName,
	)

	// Ejecutar la consulta
	row := db.QueryRow(query, id)

	if err := e.scanRows(e.model, row, e.selectColumns); err != nil {
		return err
	}

	return nil
}
