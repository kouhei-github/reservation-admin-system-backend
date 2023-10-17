package course

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type AvailableDays struct {
	value string
}

func NewAvailableDays(
	AvailableSun string,
	AvailableMon string,
	AvailableTue string,
	AvailableWed string,
	AvailableThu string,
	AvailableFri string,
	AvailableSat string,
) (*AvailableDays, error) {
	var array []string
	fromFrontArray := []string{
		AvailableSun,
		AvailableMon,
		AvailableTue,
		AvailableWed,
		AvailableThu,
		AvailableFri,
		AvailableSat,
	}
	if !slices.Contains(fromFrontArray, "true") && !slices.Contains(fromFrontArray, "false") {
		return nil, fmt.Errorf("is not right isDrinkCourse format")
	}
	for i, v := range fromFrontArray {
		if v == "true" {
			array = append(array, strconv.Itoa(i))
		}
	}
	value := strings.Join(array[:], ",")

	return &AvailableDays{value: value}, nil

}

func (c AvailableDays) ToFrontValue() map[string]string {
	returnList := map[string]string{
		"availableSun": "false",
		"availableMon": "false",
		"availableTue": "false",
		"availableWed": "false",
		"availableThu": "false",
		"availableFri": "false",
		"availableSat": "false",
	}
	FreeDrinkArray := strings.Split(c.value, ",")
	for _, v := range FreeDrinkArray {
		kbn, _ := strconv.Atoi(v)
		switch kbn {
		case 1:
			returnList["availableSun"] = "true"
		case 2:
			returnList["availableMon"] = "true"
		case 3:
			returnList["availableTue"] = "true"
		case 4:
			returnList["availableWed"] = "true"
		case 5:
			returnList["availableThu"] = "true"
		case 6:
			returnList["availableFri"] = "true"
		case 7:
			returnList["availableSat"] = "true"
		}
	}
	return returnList
}
