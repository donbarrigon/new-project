package formatter

import (
	"regexp"
	"strings"
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

// Funciones auxiliares permanecen igual
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
