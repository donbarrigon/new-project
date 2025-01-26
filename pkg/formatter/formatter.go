package formatter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var IrregularPlurals = map[string]string{
	"person":     "people",
	"child":      "children",
	"foot":       "feet",
	"tooth":      "teeth",
	"mouse":      "mice",
	"man":        "men",
	"woman":      "women",
	"ox":         "oxen",
	"cactus":     "cacti",
	"focus":      "foci",
	"analysis":   "analyses",
	"thesis":     "theses",
	"crisis":     "crises",
	"diagnosis":  "diagnoses",
	"appendix":   "appendices",
	"vertex":     "vertices",
	"index":      "indices",
	"matrix":     "matrices",
	"axis":       "axes",
	"basis":      "bases",
	"fungus":     "fungi",
	"radius":     "radii",
	"alumnus":    "alumni",
	"curriculum": "curricula",
	"datum":      "data",
	"medium":     "media",
	"forum":      "fora",
	"bacterium":  "bacteria",
	"syllabus":   "syllabi",
	"criterion":  "criteria",
	"aquarium":   "aquaria",
	"stadium":    "stadia",
	"stimulus":   "stimuli",
	"die":        "dice",
	"formula":    "formulae",
	"genus":      "genera",
	"bison":      "bison",    // no cambia
	"deer":       "deer",     // no cambia
	"sheep":      "sheep",    // no cambia
	"salmon":     "salmon",   // no cambia
	"aircraft":   "aircraft", // no cambia
	"series":     "series",   // no cambia
	"species":    "species",  // no cambia
	"fish":       "fish",     // no cambia
	"trousers":   "trousers", // no cambia
	"scissors":   "scissors", // no cambia
	"clothes":    "clothes",  // no cambia
	"news":       "news",     // no cambia
}

// ToTableName convierte el nombre a un nombre de tabla en formato snake_case y pluralizado.
func ToTableName(n string) string {

	snakeCase := ToSnakeCase(n)
	// Pluralizar
	return Pluralize(snakeCase)
}

// ToSnakeCase convierte la cadena a una cadena en formato snake_case
func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Insertar "_" antes de las letras mayúsculas y manejar espacios
	var result strings.Builder
	for i, char := range s {
		if char == ' ' {
			result.WriteRune('_')
			continue
		}

		if i > 0 && isUpper(char) &&
			((i+1 < len(s) && isLower(rune(s[i+1]))) || isLower(rune(s[i-1]))) {
			result.WriteRune('_')
		}
		result.WriteRune(char)
	}

	// Convertir a minúsculas
	snakeCase := strings.ToLower(result.String())

	// Manejar múltiples "_" consecutivos
	re := regexp.MustCompile(`_+`)
	return re.ReplaceAllString(snakeCase, "_")
}

// ToPascalCase convierte la cadena a una cadena en formato PascalCase
func ToPascalCase(s string) string {
	// Si el string está vacío, retornamos vacío
	if len(s) == 0 {
		return s
	}

	// Convertir el string a una slice de runes para manejar caracteres Unicode
	runes := []rune(s)
	var result strings.Builder

	// Variable para controlar si el siguiente carácter debe ser mayúscula
	nextUpper := true

	for i := 0; i < len(runes); i++ {
		// Si es un carácter especial o espacio, marcamos que el siguiente debe ser mayúscula
		if !unicode.IsLetter(runes[i]) && !unicode.IsNumber(runes[i]) {
			nextUpper = true
			continue
		}

		// Si es un carácter en mayúscula precedido por una letra minúscula
		// (caso camelCase), lo tratamos como el inicio de una nueva palabra
		if i > 0 && unicode.IsUpper(runes[i]) && unicode.IsLower(runes[i-1]) {
			nextUpper = true
		}

		if nextUpper {
			// Convertir a mayúscula si es necesario
			result.WriteRune(unicode.ToUpper(runes[i]))
			nextUpper = false
		} else {
			// Convertir a minúscula para el resto de los caracteres
			result.WriteRune(unicode.ToLower(runes[i]))
		}
	}

	return result.String()
}

// Funciones auxiliares
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiou", r)
}

// Pluralize pluraliza una cadena según las reglas estándar de pluralización en inglés.
func Pluralize(word string) string {
	if word == "" {
		return ""
	}

	if plural, exists := IrregularPlurals[word]; exists {
		return plural
	}

	if strings.HasSuffix(word, "y") {
		if len(word) > 1 && isVowel(rune(word[len(word)-2])) {
			return word + "s"
		}
		return word[:len(word)-1] + "ies"
	}

	if strings.HasSuffix(word, "s") ||
		strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "z") ||
		strings.HasSuffix(word, "ch") ||
		strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	if strings.HasSuffix(word, "f") {
		return word[:len(word)-1] + "ves"
	}
	if strings.HasSuffix(word, "fe") {
		return word[:len(word)-2] + "ves"
	}

	return word + "s"
}

// ToFloat64 convierte un valor a float64
func ToFloat64[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | time.Time | string](value T) (float64, error) {
	switch v := any(value).(type) {
	case int, int8, int16, int32, int64:
		return float64(v.(int64)), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(v.(uint64)), nil
	case float32, float64:
		return float64(v.(float64)), nil
	case string:
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("error al convertir string a float64: %v", err)
		}
		return num, nil
	default:
		return 0, fmt.Errorf("tipo no compatible para float64: %T", value)
	}
}

// ToInt64 convierte un valor a int64
func ToInt64[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | time.Time | string](value T) (int64, error) {
	switch v := any(value).(type) {
	case int, int8, int16, int32, int64:
		return int64(v.(int64)), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(v.(uint64)), nil
	case string:
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("error al convertir string a int64: %v", err)
		}
		return num, nil
	default:
		return 0, fmt.Errorf("tipo no compatible para int64: %T", value)
	}
}
