package Groupie

import (
	"strings"
	"time"
	"unicode"
)

func Date(date string) string {
	// 02, 01, 2006 -> 02-01-2006
	date = strings.ReplaceAll(date, ", ", "-")
	// Parse the date string using the "DD-MM-YYYY" layout
	t, err := time.Parse("02-01-2006", date)
	if err != nil {
		return date
	}
	// Format the time using the "Month D, YYYY" layout
	return t.Format("January 2, 2006")
}

func Title(s string) string {
	chars := []rune(s)

	for i, r := range chars {
		if i == 0 || chars[i-1] == ' ' || chars[i-1] == '"' {
			chars[i] = unicode.ToTitle(r)
		}
	}
	// Usa -> USA
	abbreviations := []string{"Usa", "Uk"}
	for _, e := range abbreviations {
		if strings.Contains(string(chars), e) {
			chars = []rune(strings.ReplaceAll(string(chars), e, strings.ToUpper(e)))
		}
	}

	return string(chars)
}
