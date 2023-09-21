package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3))    // Output: "III"
	fmt.Println(intToRoman(58))   // Output: "LVIII"
	fmt.Println(intToRoman(1994)) // Output: "MCMXCIV"
}

type RomanNumeral struct {
	Symbol string
	Value  int
}

var numerals = []RomanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}

	var result strings.Builder

	for _, numeral := range numerals {
		for num >= numeral.Value {
			num -= numeral.Value
			result.WriteString(numeral.Symbol)
		}
	}

	return result.String()
}
