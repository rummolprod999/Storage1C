package main

import "strings"

func ReplaceBadSymbols(s string) string {
	x := strings.ReplaceAll(s, "\u0000", "")
	return x
}
