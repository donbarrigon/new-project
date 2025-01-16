package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// Definir las reglas de validación
type Validation struct {
	Required bool
	Min      *float64
	Max      *float64
	Email    bool
	Custom   func(value any) error
}

// Función auxiliar para crear punteros a float
func FloatPtr(f float64) *float64 {
	return &f
}

// Validar si un campo es requerido
func Required(value any) error {
	if value == "" {
		return errors.New("field is required")
	}
	return nil
}

// Validar si el valor es mayor o igual al mínimo
func Min(value any, min float64) error {
	// Validar Min
	switch v := value.(type) {
	case int:
		if float64(v) < min {
			return fmt.Errorf("value must be greater than or equal to %f", min)
		}
	case float64:
		if v < min {
			return fmt.Errorf("value must be greater than or equal to %f", min)
		}
	case string:
		if len(v) < int(min) { // Para comparar la longitud con min flotante
			return fmt.Errorf("value length must be greater than or equal to %f characters", min)
		}
	default:
		return fmt.Errorf("unsupported value type for Min validation")
	}
	return nil
}

func Max(value any, max float64) error {
	// Validar Max
	switch v := value.(type) {
	case int:
		if float64(v) > max {
			return fmt.Errorf("value must be less than or equal to %f", max)
		}
	case float64:
		if v > max {
			return fmt.Errorf("value must be less than or equal to %f", max)
		}
	case string:
		if len(v) > int(max) { // Para comparar la longitud con max flotante
			return fmt.Errorf("value length must be less than or equal to %f characters", max)
		}
	default:
		return fmt.Errorf("unsupported value type for Max validation")
	}
	return nil
}

// Validar si un campo tiene un formato de correo electrónico válido
func Email(value any) error {
	if v, ok := value.(string); ok {
		// Expresión regular para validar el formato del email
		regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, _ := regexp.MatchString(regex, v)
		if !matched {
			return errors.New("invalid email format")
		}
	}
	return nil
}

// Función para validar los campos de un struct
func Struct(v any, rules map[string]Validation) error {
	// Obtener el tipo de la estructura para iterar sobre sus campos
	structValue := reflect.ValueOf(v)
	structType := reflect.TypeOf(v)

	// Iterar sobre los campos del struct
	numFields := structValue.NumField()
	for i := 0; i < numFields; i++ {
		field := structValue.Field(i)
		fieldName := structType.Field(i).Name
		fieldValue := field.Interface()

		// Obtener las reglas de validación para el campo actual
		fieldRules, exists := rules[fieldName]
		if !exists {
			continue
		}

		// Validar "required"
		if fieldRules.Required {
			if err := Required(fieldValue); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		}

		// Validar "min"
		if fieldRules.Min != nil {
			if err := Min(fieldValue, *fieldRules.Min); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		}

		// Validar "max"
		if fieldRules.Max != nil {
			if err := Max(fieldValue, *fieldRules.Max); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		}

		// Validar "email"
		if fieldRules.Email {
			if err := Email(fieldValue); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		}

		// Validación personalizada
		if fieldRules.Custom != nil {
			if err := fieldRules.Custom(fieldValue); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		}
	}
	return nil
}
