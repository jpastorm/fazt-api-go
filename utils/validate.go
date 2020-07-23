package utils

import (
	"fmt"
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func IsInteger(number string) bool {
	fmt.Println(digitCheck.MatchString(number))
	return digitCheck.MatchString(number)

}
