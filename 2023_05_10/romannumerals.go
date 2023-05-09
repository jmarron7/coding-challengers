package main

import (
	"errors"
	"log"
	"strings"
)

var romanNumerals = []struct {
	Value  int
	Symbol string
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

func intToRoman(num int) (string, error) {
	if num < 1 || num > 3999 {
		return "", errors.New("invalid input")
	}

	sb := strings.Builder{}

	for _, r := range romanNumerals {
		for num >= r.Value {
			sb.WriteString(r.Symbol)
			num -= r.Value
		}
	}

	return sb.String(), nil
}

func main() {
	num := 1469
	roman, err := intToRoman(num)
	if err != nil {
		log.Fatal(err)
	}

	println(num, "in Roman numerals is", roman)
}
