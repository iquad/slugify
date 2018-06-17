package slugify

import (
	"strings"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// replaceSet variable used for initial configuration
var replaceSet []string

// init set pre-defined configuration as default configuration
func init() {
	replaceSet = []string{
		" ", "-",
		"'", "",
		"ı", "i",
		",", "",
		".", "",
		"#", "",
		"!", "",
	}
}

// Slugify converts given string into unicode compatible
// slug with standart Turkish configuration
func Slugify(str string) (result string, err error) {
	replacer := strings.NewReplacer(replaceSet...)
	return transformizer(replacer, str)
}

// transformizer replaces unicode string with their ASCII equals
func transformizer(replacer *strings.Replacer, str string) (result string, err error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	preResult, _, err := transform.String(t, str)

	if err != nil {
		return result, err
	}

	result = strings.ToLower(replacer.Replace(preResult))
	return
}
