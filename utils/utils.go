package utils

import (
	"errors"
	"fmt"
	"strings"
)

// MinutesFromHMFormat will return the minutes in common 01h03m
// human readable format
func MinutesFromHMFormat(input string) (int, error) {
	if len(input) == 0 {
		return 0, errors.New("incorrect format. emtpy string")
	}

	if strings.Count(input, "m") != 1 {
		return 0, fmt.Errorf("incorrect format: %s. please use 1h23m", input)
	}

	if strings.Count(input, "h") > 1 {
		return 0, fmt.Errorf("incorrect format: %s. please use 1h23m", input)
	}

	lastCharacter := input[len(input)-1:]
	if lastCharacter != "m" {
		return 0, fmt.Errorf("incorrect format: %s. please use 1h23m", input)
	}

	sign := 1
	temp := 0
	minutes := 0

format_parser:
	for _, ch := range input {
		switch ch {
		case '+':
			sign = +1
		case '-':
			sign = -1
		case 'm':
			minutes += temp
			break format_parser
		case 'h':
			minutes += temp * 60
			temp = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			temp = temp*10 + int(ch-'0')
		default:
			return 0, fmt.Errorf("incorrect character: %c. please use 1h23m", ch)
		}
	}

	if minutes > 0 {
		return sign * minutes, nil
	} else {
		return 0, errors.New("incorrect format: zero time. please use 1h23m")
	}
}
