package request

type FormRequest interface {
	Rules() FormRequest
	PrepareForValidation() error //Propósito: Modifica o normaliza los datos del request antes de validar.
	WithValidator() error        // Propósito: Permite añadir lógica adicional después de preparar el validador pero antes de que se realice la validación.
}
