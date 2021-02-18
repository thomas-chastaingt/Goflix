package utils

import (
	"regexp"
)

//DeleteSpecialCharacter should protect from SQL Injection and XSS
func DeleteSpecialCharacter(text string) string {
	reg := regexp.MustCompile(`[^0-9a-zA-Z]+`)
	res := reg.ReplaceAllString(text, "")
	return res
}
