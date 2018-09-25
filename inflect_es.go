package inflect

// NewSpanishRuleset creates a simplified version of the grammar rules
// probably enough for database tables names that are written in spanish without tonic accents
// For an exhaustive ruleset you would need to detect the word's toinic sylables
// https://espanol.lingolia.com/es/gramatica/sustantivos-articulos/plural
func NewSpanishRuleset() *Ruleset {
	rs := NewRuleset()

	rs.AddPlural("í", "íes")
	rs.AddPlural("ú", "úes")
	rs.AddPlural("d", "des")
	rs.AddPlural("j", "jes")
	rs.AddPlural("l", "les")
	rs.AddPlural("n", "nes")
	rs.AddPlural("r", "res")
	rs.AddPlural("z", "ces")

	rs.AddPlural("a", "as")
	rs.AddPlural("e", "es")
	rs.AddPlural("i", "is")
	rs.AddPlural("o", "os")
	rs.AddPlural("u", "us")

	//

	rs.AddSingular("us", "u")
	rs.AddSingular("os", "o")
	rs.AddSingular("is", "i")
	rs.AddSingular("es", "e")
	rs.AddSingular("as", "a")

	rs.AddSingular("ces", "z")
	rs.AddSingular("res", "r")
	rs.AddSingular("nes", "n")
	rs.AddSingular("les", "l")
	rs.AddSingular("jes", "j")
	rs.AddSingular("des", "d")
	rs.AddSingular("úes", "ú")
	rs.AddSingular("íes", "í")

	//english consonnat ending words

	return rs
}
