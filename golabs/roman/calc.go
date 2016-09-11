
/*
Package roman provides calculators for converting an arabic number
to a roman glyph and vice versa
*/
package roman

type tuple struct {
        limit int
        glyph string
}

var romanValues = func(tuples [13]tuple) map[rune]int {
        vals := map[rune]int{}
        for _, t := range tuples {
                if len(t.glyph) == 1 {
                        vals[rune(t.glyph[0])] = t.limit
                }
        }
        return vals
}(RomanGlyphs)

// RomanGlyphs provides a conversion table arabic->roman
var RomanGlyphs = [...]tuple{
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

// ToRoman converts an arabic number to a Roman glyph
func ToRoman(n int) string {
        var res string
        for _, tuple := range RomanGlyphs {
                for n >= tuple.limit {
                        res += tuple.glyph
                        n -= tuple.limit
                }

        }
        return res
}

// ToArabic converts a roman glyph to arabic
func ToArabic(s string) int {
        var res = 0
        for index, r := range s {
                if index+1 < len(s) && romanValues[r] < romanValues[[]rune(s)[index+1]] {
                        res -= romanValues[r]
                } else {
                        res += romanValues[r]
                }
        }
        return res
}

