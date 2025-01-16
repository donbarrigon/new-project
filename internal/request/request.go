package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"

	"github.com/erespereza/clan-de-raid/pkg/validation"
)

type FormRequest interface {
	Rules() map[string]validation.Validation // proposito retornar las reglas de validacion
	PrepareForValidation() error             // Propósito: Modifica o normaliza los datos del request y añadir lógica adicional antes de validar.
	WithValidator() error                    // Propósito: Permite añadir lógica adicional después de preparar el validador pero antes de que se realice la validación.
}

func Validate(request FormRequest, r *http.Request) error {

	// Usar reflect para validar que se trabaja con el tipo especifico y obtener el tipo de request y deserializar en el tipo real
	requestValue := reflect.ValueOf(request)
	if requestValue.Kind() != reflect.Ptr || requestValue.IsNil() {
		return errors.New("se espera un puntero al tipo que implementa FormRequest")
	}

	// Leer el cuerpo de la solicitud
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// Deserializar el JSON en el struct
	if err := json.Unmarshal(body, request); err != nil {
		return err
	}

	// Preparar el request antes de validar
	if err := request.PrepareForValidation(); err != nil {
		return err
	}

	// Añadir lógica adicional después de preparar el validador
	if err := request.WithValidator(); err != nil {
		return err
	}

	// Validar el request con las reglas de validación
	if err := validation.Struct(request, request.Rules()); err != nil {
		return err
	}

	return nil
}
