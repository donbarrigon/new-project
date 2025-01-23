package request

// Definir el struct a validar
type User struct {
	Name  string `json:"name" rules:"required|min:3|max:10"`
	Email string `json:"email" rules:"required|email"`
	Age   int    `json:"age" rules:"min:18"`
}

func (u *User) PrepareForValidation() error {
	// Añadir lógica adicional para normalizar los datos del request (por ejemplo, reemplazar espacios en blanco con guiones bajos)
	return nil
}

func (u *User) WithValidator() error {
	// Añadir lógica adicional después de preparar el validador (por ejemplo, restringir el uso de palabras ofensivas)
	return nil
}
