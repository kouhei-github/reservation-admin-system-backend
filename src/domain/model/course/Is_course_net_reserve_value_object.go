package course

import "fmt"

type IsCourseNetReserve struct {
	value bool
}

func NewIsCourseNetReserve(fromFrontValue string) (*IsCourseNetReserve, error) {
	var value bool
	switch fromFrontValue {
	case "対応":
		value = true
	case "非対応":
		value = false
	default:
		return nil, fmt.Errorf("is not right isCourseNetReserve format")
	}
	return &IsCourseNetReserve{value: value}, nil
}

func (c IsCourseNetReserve) ToFrontValue() string {
	var value string
	switch c.value {
	case true:
		value = "対応"
	case false:
		value = "非対応"
	}
	return value
}
