
package roman

import (
	"errors"
	"fmt"
)

// RomanNumeral represents an integer value that can be converted to Roman numerals
type RomanNumeral int

// ErrOutOfRange is returned when trying to create a RomanNumeral with an invalid value
var ErrOutOfRange = errors.New("value must be between 0 and 3999")

// NewRomanNumeral creates a new RomanNumeral with validation
func NewRomanNumeral(value int) (RomanNumeral, error) {
	if value < 0 || value > 3999 {
		return 0, ErrOutOfRange
	}
	return RomanNumeral(value), nil
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
	
	// Roman numeral conversion logic would go here
	// This is a placeholder for the actual conversion implementation
	return fmt.Sprintf("Roman(%d)", int(r))
}
