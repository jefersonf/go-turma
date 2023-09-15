package turma

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func splitFullname(s string) ([]string, error) {
	p := strings.Split(strings.Trim(s, "\"'-_"), " ")
	if len(p) == 0 {
		return []string{}, errors.New("empty name")
	}
	return p, nil
}

// TODO move this func to another place
func GetFirstAndLastName(fullname string) (string, error) {
	p, err := splitFullname(fullname)
	if err != nil {
		return "", fmt.Errorf("unable to get first name: %w", err)
	}
	firstName := Capitalize(p[0])
	if len(p) > 1 {
		lastName := Capitalize(p[len(p)-1])
		return fmt.Sprintf("%s %s", firstName, lastName), nil
	}
	return firstName, nil
}

func Capitalize(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
