package roman

import (
  "fmt"
  "math/rand"
  "reflect"
  "testing"
  "testing/quick"
)

func TestRomanNumerals(t *testing.T) {
  cases := []struct {
    Arabic int
    Roman  string
  }{
    {Arabic: 1, Roman: "I"},
    {Arabic: 2, Roman: "II"},
    {Arabic: 3, Roman: "III"},
    {Arabic: 4, Roman: "IV"},
    {Arabic: 5, Roman: "V"},
    {Arabic: 6, Roman: "VI"},
    {Arabic: 7, Roman: "VII"},
    {Arabic: 8, Roman: "VIII"},
    {Arabic: 9, Roman: "IX"},
    {Arabic: 10, Roman: "X"},
    {Arabic: 14, Roman: "XIV"},
    {Arabic: 18, Roman: "XVIII"},
    {Arabic: 20, Roman: "XX"},
    {Arabic: 39, Roman: "XXXIX"},
    {Arabic: 40, Roman: "XL"},
    {Arabic: 47, Roman: "XLVII"},
    {Arabic: 49, Roman: "XLIX"},
    {Arabic: 50, Roman: "L"},
    {Arabic: 100, Roman: "C"},
    {Arabic: 90, Roman: "XC"},
    {Arabic: 400, Roman: "CD"},
    {Arabic: 500, Roman: "D"},
    {Arabic: 900, Roman: "CM"},
    {Arabic: 1000, Roman: "M"},
    {Arabic: 1984, Roman: "MCMLXXXIV"},
    {Arabic: 3999, Roman: "MMMCMXCIX"},
    {Arabic: 2014, Roman: "MMXIV"},
    {Arabic: 1006, Roman: "MVI"},
    {Arabic: 798, Roman: "DCCXCVIII"},
  }
  for _, test := range cases {
    t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
      got := ConvertToRoman(test.Arabic)
      if got != test.Roman {
        t.Errorf("got %q, want %q", got, test.Roman)
      }
    })
  }
}

func TestPropertiesOfConversion(t *testing.T) {
  assertion := func(arabic int) bool {
    roman := ConvertToRoman(arabic)
    fromRoman := ConvertToArabic(roman)
    return fromRoman == arabic
  }

  config := &quick.Config{
    MaxCount: 1000,
    Values: func(values []reflect.Value, rand *rand.Rand) {
      // Generate random integer between 0-3999
      values[0] = reflect.ValueOf(rand.Intn(4000))
    },
  }

  if err := quick.Check(assertion, config); err != nil {
    t.Error("failed checks", err)
  }
}

func TestBruteForce(t *testing.T) {
  for i := 0; i < 4000; i++ {
    roman := ConvertToRoman(i)
    fromRoman := ConvertToArabic(roman)
    if fromRoman != i {
      t.Errorf("got %d, want %d", fromRoman, i)
    }
  }
}