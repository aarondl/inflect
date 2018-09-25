package inflect

import (
	"testing"
)

var SingularToPluralES = map[string]string{
	"libro":     "libros",
	"mesa":      "mesas",
	"jabalí":    "jabalíes",
	"bambú":     "bambúes",
	"ordenador": "ordenadores",
	"reloj":     "relojes",
	"cruz":      "cruces",
	"elefante":  "elefantes",
	"manzana":   "manzanas",
	//"club":      "clubes", this fails under our simlpe model
}

var ES = NewSpanishRuleset()

func TestPluralizeSingularES(t *testing.T) {
	for singular, plural := range SingularToPluralES {
		assertEqual(t, plural, ES.Pluralize(singular))
		assertEqual(t, ES.Capitalize(plural), ES.Capitalize(ES.Pluralize(singular)))
	}
}

func TestSingularizePluralES(t *testing.T) {
	for singular, plural := range SingularToPluralES {
		assertEqual(t, singular, ES.Singularize(plural))
		assertEqual(t, ES.Capitalize(singular), ES.Capitalize(ES.Singularize(plural)))
	}
}

func TestPluralizePluralES(t *testing.T) {
	for _, plural := range SingularToPluralES {
		assertEqual(t, plural, ES.Pluralize(plural))
		assertEqual(t, ES.Capitalize(plural), ES.Capitalize(ES.Pluralize(plural)))
	}
}

func TestSingularizeSingularES(t *testing.T) {
	for singular := range SingularToPluralES {
		assertEqual(t, singular, ES.Singularize(singular))
		assertEqual(t, ES.Capitalize(singular), ES.Capitalize(ES.Singularize(singular)))
	}
}
