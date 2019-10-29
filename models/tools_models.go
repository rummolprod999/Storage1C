package models

import (
	"io"
	"strings"
)

var Debug = false

func ReplaceBadSymbols(s string) string {
	x := strings.ReplaceAll(s, "\u0000", "")
	return x
}
func identReader(encoding string, input io.Reader) (io.Reader, error) {
	return input, nil
}
