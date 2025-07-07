
package roman

import (
	"fmt"
	"strings"
	)

func ConvertToRoman(n int) string {
	rn := NewRomanNumeral(n)
	return rn.ToRoman()
}

type RND struct {
	Value  int
	Symbol string
}

// earlier..
var allRomanNumerals = []RND{
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

// later..
func ConvertToArabic(roman string) int {
	var arabic = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}

// RomanNumeral represents an integer value that can be converted to Roman numerals
type RomanNumeral int

// NewRomanNumeral creates a new RomanNumeral with validation
func NewRomanNumeral(value int) RomanNumeral {
	if value < 0  {
		return RomanNumeral(0)
	}
	if value > 3999 {
		return RomanNumeral(3999)
	}
	
	return RomanNumeral(value)
}

// Value returns the integer value of the RomanNumeral
func (r RomanNumeral) Value() int {
	return int(r)
}

// String returns the string representation of the integer value
func (r RomanNumeral) String() string {
	return fmt.Sprintf("%d", int(r))
}

// ToRoman converts the RomanNumeral to its Roman numeral string representation
func (r RomanNumeral) ToRoman() string {
	if r == 0 {
		return ""
	}
	
	// Mapping of values to Roman numerals in descending order
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	
	result := ""
	num := int(r)
	
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			result += numerals[i]
			num -= values[i]
		}
	}
	
	return result
}
