package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/erespereza/new-project/internal/orm"
	"github.com/erespereza/new-project/internal/request"
	"github.com/erespereza/new-project/pkg/utils"
)

type ControllerFunc func(ctx *Context)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Body    any
	User    *orm.Model
	Errors  map[string]string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  w,
	}
}

// Toma los valores de la url y los parsea en un map y lo retorna
func (c *Context) ParseQuery() map[string]any {

	// Inicializar el mapa Query si no está inicializado
	query := make(map[string]any)

	// Obtener los parámetros de la URL
	queryParams := c.Request.URL.Query()

	// Iterar sobre los parámetros de la URL
	for key, values := range queryParams {
		// El valor puede ser un solo valor o una lista, tomo solo el primer valor ya mas adelante cuando tenga tiempo me molesto en hacerlo con una lista ya que no se si eso es comun
		value := ""
		if len(values) > 0 {
			value = values[0]
		}

		// Intentar convertir el valor a diferentes tipos
		if intValue, err := strconv.Atoi(value); err == nil {
			// Es un int
			query[key] = intValue
		} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			// Es un float
			query[key] = floatValue
		} else if boolValue, err := strconv.ParseBool(value); err == nil {
			// Es un bool
			query[key] = boolValue
		} else {
			// Es un string (por defecto)
			query[key] = value
		}
	}
	return query
}

// Toma los valores de la cabecera y los retorna manejables
func (c *Context) Query(key ...string) any {
	//validar para evitar un error
	if c.Request == nil {
		return nil
	}

	// Obtener los parámetros de la URL para buscar el valor deseado
	if len(key) == 0 {
		// retorna todos los parametros de la query en un map
		return c.queryToMap()
	} else if len(key) == 1 {
		// Return el valor para una sola key
		return c.queryParam(key[0])
	} else {
		// retorna los valores para multeples keys
		result := make(map[string]string)
		for _, k := range key {
			result[k] = c.queryParam(k)
		}
		return result
	}
}

// queryToMap retorna todos los parametros de la query en un map[string]string
func (c *Context) queryToMap() map[string]string {
	query := c.Request.URL.Query()
	result := make(map[string]string)
	for k, v := range query {
		if len(v) > 0 {
			result[k] = v[0]
		}
	}
	return result
}

// queryParam retorna el valor para una sola key
func (c *Context) queryParam(key string) string {
	values := c.Request.URL.Query()[key]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

// Validate valida los datos del request y los guarda en c.Body
// forma de uso if err := ctx.Validate([]FormRequest{}) err != nil { return err }
func (c *Context) Validate(req request.FormRequest) error {

	isSlice, err1 := c.getBody(&req)
	if err1 != nil {
		return err1
	}

	if isSlice {
		for _, r := range *c.Body.(*[]request.FormRequest) {
			err := validateStruct(&r)
			if err != nil {
				return err
			}
		}
	} else {
		if err := validateStruct(c.Body.(*request.FormRequest)); err != nil {
			return err
		}
	}

	return nil
}

func validateStruct(req *request.FormRequest) error {
	fmt.Println(req)
	return nil
}

func (c *Context) getBody(req *request.FormRequest) (bool, error) {
	// Leer completamente el cuerpo de la solicitud
	isSlice := false
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return isSlice, fmt.Errorf("error leyendo el cuerpo de la solicitud: %v", err)
	}
	defer c.Request.Body.Close()

	// Eliminar espacios en blanco
	body = bytes.TrimSpace(body)

	switch body[0] {
	case '[':
		// JSON array - slice de implementaciones de FormRequest
		sliceReq := utils.StructToSlice(*req)
		if err := json.Unmarshal(body, &sliceReq); err != nil {
			return isSlice, fmt.Errorf("error decodificando array JSON: %v", err)
		}
		c.Body = &sliceReq
		isSlice = true
	case '{':
		// JSON objeto
		if err := json.Unmarshal(body, req); err != nil {
			return isSlice, fmt.Errorf("error decodificando objeto JSON: %v", err)
		}
		c.Body = req
	default:
		return isSlice, errors.New("el JSON debe ser un array u objeto")
	}

	return isSlice, nil
}

// {
//     "status": "error",
//     "message": "Validation failed",
//     "code": 422,
//     "errors": {
//         "email": [
//             "The email field is required.",
//             "The email must be a valid email address."
//         ],
//         "password": [
//             "The password field is required.",
//             "The password must be at least 8 characters."
//         ],
//         "username": [
//             "The username field is required."
//         ]
//     }
// }

// {
//     "status": "error",
//     "message": "Error processing collection",
//     "code": 422,
//     "errors": [
//         {
//             "index": 0,
//             "errors": {
//                 "name": ["The name field is required."],
//                 "email": ["The email must be a valid email address."]
//             }
//         },
//         {
//             "index": 1,
//             "errors": {
//                 "name": ["The name field is required."]
//             }
//         }
//     ]
// }

/*
func (c *Context) Validate_descartado(req *request.FormRequest) error {
	// decodificar el body
	decoder := json.NewDecoder(c.Request.Body)

	//tomar el token para sabaer si es array o objeto
	token, err := decoder.Token()
	if err != nil {
		log.Fatal("Error leyendo el token:", err)
	}

	switch delim := token.(type) {
	case json.Delim:
		if delim == '[' {
			// el json es un array
			// Procesar el array
			var elemen []request.FormRequest
			for decoder.More() {

				if err := decoder.Decode(req); err != nil {
					return fmt.Errorf("Error decodificando elemento del array: %v", err)
				}
				elemen = append(elemen, *req)
			}
			c.Body = elemen
		} else if delim == '{' {
			// el json es un objeto
			// Procesar el objeto y guardarlo en c.Body
			if err := decoder.Decode(req); err != nil {
				return fmt.Errorf("error decodificando el objeto: %v", err)
			}
			c.Body = *req
		}
		return nil
	default:
		return errors.New("El JSON no es un array ni un objeto valido")
	}
}
*/
