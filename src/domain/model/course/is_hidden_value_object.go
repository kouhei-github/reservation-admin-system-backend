package course

import "fmt"

type IsHidden struct {
	value bool
}

func NewIsHidden(fromFrontValue string) (*IsHidden, error) {
	var value bool
	switch fromFrontValue {
	case "掲載":
		value = false
	case "非掲載":
		value = true
	default:
		return nil, fmt.Errorf("is not right isHidden format")
	}
	return &IsHidden{value: value}, nil
}

func (c IsHidden) ToFrontValue() string {
	var value string
	switch c.value {
	case true:
		value = "掲載"
	case false:
		value = "非掲載"
	}
	return value

}
