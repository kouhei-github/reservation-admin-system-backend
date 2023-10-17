package course

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type TargetValue struct {
	value string
}

func CheckCharLength(checkValue string, checkLength int) (*TargetValue, error) {
	value := checkValue
	if utf8.RuneCountInString(value) > checkLength {
		return nil, fmt.Errorf("is invalid number of characters (" + strconv.Itoa(checkLength) + ")")
	}
	return &TargetValue{value: value}, nil
}
