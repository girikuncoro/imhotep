
package roman_test

import (
        "testing"

        "github.com/derailed/imhotep/golabs/roman"
)

var useCases = []struct {
        arabic int
        glyph  string
}{
        {1, "I"},
        {2, "II"},
        {4, "IV"},
        {5, "V"},
        {6, "VI"},
        {9, "IX"},
        {10, "X"},
        {11, "XI"},
        {20, "XX"},
        {21, "XXI"},
        {25, "XXV"},
        {30, "XXX"},
        {35, "XXXV"},
        {40, "XL"},
        {50, "L"},
        {90, "XC"},
        {100, "C"},
        {400, "CD"},
        {500, "D"},
        {900, "CM"},
        {1000, "M"},
        {2016, "MMXVI"},
        {3999, "MMMCMXCIX"},
        {0, ""},
}

func TestArabicToRoman(t *testing.T) {
        for _, uc := range useCases {
                actual := roman.ToRoman(uc.arabic)
                expected := uc.glyph

                if actual != expected {
                        t.Fatalf("(%d) Expected %s GOT %s", uc.arabic, expected, actual)
                }
        }
}

func TestRomanToArabic(t *testing.T) {
        for _, uc := range useCases {
                if actual := roman.ToArabic(uc.glyph); actual != uc.arabic {
                        t.Fatalf("Roman For %s -- Got %d expected `%d`", uc.glyph, actual, uc.arabic)
                }
        }
}

