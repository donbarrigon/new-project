package request

import (
	"github.com/erespereza/clan-de-raid/pkg/validation"
)

// Definir el struct a validar
type User struct {
	Request
	Name  string
	Email string
	Age   int
}

// Define las reglas para la validacion
func (u *User) Rules() map[string]validation.Validation {
	return map[string]validation.Validation{
		"Name":  {Required: true},
		"Email": {Required: true, Email: true},
		"Age":   {Min: validation.FloatPtr(0), Max: validation.FloatPtr(120)},
	}
}

func (u *User) PrepareForValidation() error {
	// Añadir lógica adicional para normalizar los datos del request (por ejemplo, reemplazar espacios en blanco con guiones bajos)
	return nil
}

func (u *User) WithValidator() error {
	// Añadir lógica adicional para añadir lógica adicional después de preparar el validador (por ejemplo, restringir el máximo de caracteres en el nombre)
	return nil
}
