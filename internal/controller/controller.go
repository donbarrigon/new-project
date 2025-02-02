package controller

import (
	"net/http"
	"strconv"

	"github.com/donbarrigon/new-project/internal/orm"
)

type ControllerFunc func(ctx *Context)

type ValidationError map[string]any
type FieldsError map[string][]string

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Body    any
	// isSlice bool
	User   *orm.Model
	Errors map[string]string
}

type ErrorResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Errors     interface{} `json:"errors"`
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request: r,
		Writer:  w,
	}
}

// Toma los valores de la url y los parsea en un map[string]any o [int | float64 | bool | string] si es solo 1
func (c *Context) ParseQuery(key ...string) any {

	// Obtener los parámetros de la URL
	queryParams := c.Query(key...)

	if len(key) == 1 {
		return c.parseValue(queryParams.(string))
	}

	// Inicializar el mapa Query a retornar
	query := make(map[string]any)

	// Iterar sobre los parámetros de la URL
	for k, value := range queryParams.(map[string]string) {
		query[k] = c.parseValue(value)
	}
	return query
}

// Suggested code may be subject to a license. Learn more: ~LicenseLog:2687351136.
func (c *Context) parseValue(value string) any {
	if intValue, err := strconv.Atoi(value); err == nil {
		// Es un int
		return intValue
	} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		// Es un float
		return floatValue
	} else if boolValue, err := strconv.ParseBool(value); err == nil {
		// Es un bool
		return boolValue
	} else {
		// Es un string (por defecto)
		return value
	}
}

// Query es una forma menos verbosa de tomar los valores de la cabecera y los retorna en un map o string si es solo 1
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
// forma de uso if err := ctx.Validate(&FormRequest{}) err != nil { return err }
// func (c *Context) Validate(req *request.FormRequest) ValidationError {

// 	if err := c.getBody(req); err != nil {
// 		return nil
// 	}

// 	if c.isSlice {
// 		for _, r := range *c.Body.(*[]request.FormRequest) {
// 			err := validateStruct(&r)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	} else {
// 		if err := validateStruct(c.Body.(*request.FormRequest)); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (c *Context) getBody(req *request.FormRequest) error {

// 	c.isSlice = false

// 	// Leer completamente el cuerpo de la solicitud
// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		return fmt.Errorf("error leyendo el cuerpo de la solicitud: %v", err)
// 	}
// 	defer c.Request.Body.Close()

// 	// Eliminar espacios en blanco
// 	body = bytes.TrimSpace(body)

// 	switch body[0] {
// 	case '[':
// 		// JSON array - slice de implementaciones de FormRequest
// 		sliceReq := utils.StructToSlice(*req)
// 		if err := json.Unmarshal(body, &sliceReq); err != nil {
// 			return fmt.Errorf("error decodificando array JSON: %v", err)
// 		}
// 		c.Body = &sliceReq
// 		c.isSlice = true
// 	case '{':
// 		// JSON objeto
// 		if err := json.Unmarshal(body, req); err != nil {
// 			return fmt.Errorf("error decodificando objeto JSON: %v", err)
// 		}
// 		c.Body = req
// 	default:
// 		return errors.New("el JSON debe ser un array u objeto")
// 	}

// 	return nil
// }

// func validateStruct(req *request.FormRequest) FieldsError {
// 	errors := make(FieldsError)
// 	v := reflect.ValueOf(req)

// 	// obtener el valor al que apunta el puntero
// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}

// 	t := v.Type()

// 	numFilds := v.NumField()
// 	for i := 0; i < numFilds; i++ {
// 		field := v.Field(i)
// 		fieldType := t.Field(i)

// 		// Obtener las reglas de validación del tag
// 		rulesTag := fieldType.Tag.Get("rules")
// 		if rulesTag == "" {
// 			// si no tiene reglas de validación, pasar al siguiente campo
// 			continue
// 		}

// 		// Obtener el nombre JSON del campo
// 		jsonTag := fieldType.Tag.Get("json")
// 		fieldName := jsonTag
// 		if fieldName == "" {
// 			fieldName = formatter.ToSnakeCase(fieldType.Name)
// 		}

// 		// Procesar cada regla de validación
// 		// si las reglas estan separadas por pipe
// 		rules := strings.Split(rulesTag, "|")
// 		if len(rules) == 1 {
// 			// o por si si las reglas estan separadas por comas
// 			rules = strings.Split(rulesTag, ",")
// 		}
// 		for _, rule := range rules {
// 			validationError := validateField(field, rule)
// 			if validationError != "" {
// 				errors[fieldName] = append(errors[fieldName], validationError)
// 			}
// 		}
// 	}

// 	if len(errors) > 0 {
// 		return errors
// 	}
// 	return nil
// }

// func validateField(field reflect.Value, rule string) string {
// 	// Separar regla y parámetro si existe por :
// 	parts := strings.Split(rule, ":")
// 	if len(parts) == 1 {
// 		// Separar regla y parámetro si existe por =
// 		parts = strings.Split(rule, "=")
// 	}
// 	ruleType := parts[0]

// 	switch ruleType {
// 	case "required":
// 		if err := validation.Required(field.Interface()); err != nil {
// 			return err.Error()
// 		}
// 	case "min":
// 		if len(parts) > 1 {
// 			// paso de tener que saber cual es el tipo de dato para poder obtener su valo
// 			// me tocaria crear otra funcion que determine el tipo y tome el valor correcto
// 			// con tanta maricada que hay que hacer para esto mejor trabajo con maps
// 			// voy a dejar esto comentado y archivado por si acaso y trabajo mejor con maps.
// 			// ademas que lo mas seguro es que las ventajas de performance de usar un struct lo mas seguro es que se pierdan
// 			// con tanta maricada que hay que hacer para cualquier cosa
//			// puta de verdad que me mame de trabajar con structs te restringen tanto que hay que hacer el triple para que todo funcione
// 			if err := validation.Min(field.Interface(), parts[1]); err != nil {
// 				return err.Error()
// 			}
// 		}
// 	case "max":
// 		if len(parts) > 1 {
// 			if err := validation.Max(field.Interface(), parts[1]); err != nil {
// 				return err.Error()
// 			}
// 		}
// 	case "email":
// 		if err := validation.Email(field.Interface()); err != nil {
// 			return err.Error()
// 		}
// 	}
// 	return ""
// }

// formato para retornar errores
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
