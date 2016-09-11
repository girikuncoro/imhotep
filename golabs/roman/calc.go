/*
Package roman provides an arabic number to roman calculator
*/
package roman

// RomanGlyphs provides a conversion table arabic->roman
var RomanGlyphs = [...]struct {
	limit int
	glyph string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ArabicToRoman converts an arabic number to a Roman glyph
func ArabicToRoman(n int) string {
	var res string
	for _, tuple := range RomanGlyphs {
		for n >= tuple.limit {
			res += tuple.glyph
			n -= tuple.limit
		}

	}
	return res
}

// Fred is cool
func Fred() {}
