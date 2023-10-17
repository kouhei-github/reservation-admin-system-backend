package course

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type FreeDrinkKbn struct {
	value string
}

func NewFreeDrinkKbn(IsDrinkCourse string, IsDrinkCourseOver3h string, IsDrinkCourseOnly string) (*FreeDrinkKbn, error) {
	var array []string
	fromFrontArray := []string{IsDrinkCourse, IsDrinkCourseOver3h, IsDrinkCourseOnly}
	// エラーキャッチ
	if !slices.Contains(fromFrontArray, "true") && !slices.Contains(fromFrontArray, "false") {
		return nil, fmt.Errorf("is not right isDrinkCourse format")
	}

	for i, v := range fromFrontArray {
		if v == "true" {
			array = append(array, strconv.Itoa(i))
		}
	}
	value := strings.Join(array[:], ",")

	return &FreeDrinkKbn{value: value}, nil
}

func (c FreeDrinkKbn) ToFrontValue() map[string]string {
	returnList := map[string]string{
		"isDrinkCourse":       "false",
		"isDrinkCourseOver3h": "false",
		"isDrinkCourseOnly":   "false",
	}
	FreeDrinkArray := strings.Split(c.value, ",")
	for _, v := range FreeDrinkArray {
		kbn, _ := strconv.Atoi(v)
		switch kbn {
		case 1:
			returnList["isDrinkCourse"] = "true"
		case 2:
			returnList["isDrinkCourseOver3h"] = "true"
		case 3:
			returnList["isDrinkCourseOnly"] = "true"
		}
	}
	return returnList
}
