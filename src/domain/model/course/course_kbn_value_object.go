package course

import "fmt"

type CourseKbn struct {
	value int
}

func NewCourseKbn(fromFrontValue string) (*CourseKbn, error) {
	var value int
	switch fromFrontValue {
	case "コース料理":
		value = 1
	case "席のみ":
		value = 2
	default:
		return nil, fmt.Errorf("is not right courseKbn format")
	}
	return &CourseKbn{value: value}, nil
}

func (c CourseKbn) ToFrontValue() string {
	var value string
	switch c.value {
	case 1:
		value = "コース料理"
	case 2:
		value = "席のみ"
	}
	return value
}
