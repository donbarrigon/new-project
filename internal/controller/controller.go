package controller

import (
	"net/http"
	"strconv"

	"github.com/erespereza/new-project/internal/orm"
)

type ControllerFunc func(ctx *Context)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	User    *orm.Model
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
