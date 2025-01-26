package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/erespereza/new-project/pkg/formatter"
)

type AllowTypes interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | time.Time | string
}

// Required valida que el campo no sea nulo ni vacío.
func Required(value any) error {
	if value == nil {
		return errors.New("el campo es obligatorio")
	}

	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return errors.New("el campo es obligatorio")
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		// Números diferentes cero son válidos
		if v != 0 {
			return errors.New("el campo es obligatorio")
		}
	case float32, float64:
		// Números decimales diferentes de cero siempre válidos
		if v != 0 {
			return errors.New("el campo es obligatorio")
		}
	case bool:
		// Booleanos true son siempre válidos
		if !v {
			return errors.New("el campo es obligatorio")
		}
	case []interface{}:
		if len(v) == 0 {
			return errors.New("el campo es obligatorio")
		}
	case map[string]interface{}:
		if len(v) == 0 {
			return errors.New("el campo es obligatorio")
		}
	case time.Time:
		if v.IsZero() {
			return errors.New("el campo es obligatorio")
		}
	default:
		// Para tipos personalizados/structs, verificar si implementa IsZero() bool
		if i, ok := value.(interface{ IsZero() bool }); ok {
			if i.IsZero() {
				return errors.New("el campo es obligatorio")
			}
		}
	}

	return nil
}

// Min valida si el valor es mayor o igual al mínimo.
// Los strings se validan que la longitud de caracteres sea mayor o igual al mínimo.
func Min[T AllowTypes, U AllowTypes](value T, min U) error {
	switch v := any(value).(type) {
	case string:
		// Convierte los valores a int64 para comparar
		m, err := formatter.ToInt64(min)
		if err != nil {
			return err
		}
		// Se valida la longitud y se convierte len a int64 para que sean del mismo tipo y poder comparar
		if int64(len(v)) < m {
			return fmt.Errorf("la longitud del texto %d es menor que la longitud mínima permitida de %d", len(v), m)
		}
	default:
		// Convierte los valores a float64 para comparar
		valueFloat, err1 := formatter.ToFloat64(value)
		minFloat, err2 := formatter.ToFloat64(min)
		if err1 != nil {
			return fmt.Errorf("error al convertir el valor a float64: %v", err1)
		}
		if err2 != nil {
			return fmt.Errorf("error al convertir el valor minimo a float64: %v", err2)
		}

		// Comparación de valores numéricos
		if valueFloat < minFloat {
			return fmt.Errorf("el valor %v es menor que el valor mínimo permitido de %v", value, min)
		}
	}
	return nil
}

// Max Valida si el valor es menor o igual al máximo.
// los strings se valida que la longitud de caracteres sea menor o igual al máximo.
func Max[T AllowTypes, U AllowTypes](value T, max U) error {
	switch v := any(value).(type) {
	case string:
		// convierte los valores a int64 para comparar
		m, err := formatter.ToInt64(max)
		if err != nil {
			return err
		}
		//se valida la longitud y se conviente len a int 64 para que sean del mismo tipo y poder comparar
		if int64(len(v)) > m {
			return fmt.Errorf("la longitud del texto %d excede el maximo permitido de %d", len(v), m)
		}
	default:
		// convierte los valores a float64 para comparar
		valueFloat, err1 := formatter.ToFloat64(value)
		maxFloat, err2 := formatter.ToFloat64(max)
		if err1 != nil {
			return fmt.Errorf("error al convertir el valor a float64: %v", err1)
		}
		if err2 != nil {
			return fmt.Errorf("error al convertir el valor máximo a float64: %v", err2)
		}

		// Comparación de valores numéricos
		if valueFloat > maxFloat {
			return fmt.Errorf("el valor %v es mayor que el valor maximo permitido de %v", value, max)
		}
	}
	return nil
}

// Email valida si un campo tiene un formato de correo electrónico válido
func Email(value string) error {
	// Expresión regular para validar el formato del email
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, value)
	if !matched {
		return errors.New("invalid email format")
	}
	return nil
}
