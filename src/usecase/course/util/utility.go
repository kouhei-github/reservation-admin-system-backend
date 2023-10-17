package util

import (
	"encoding/json"
	"net-http/myapp/domain/model/course"
)

type Utility struct{}

func NewUtility() *Utility {
	return &Utility{}
}

func (u Utility) GetCourseDataToArray(courses []course.Course) ([][]string, error) {
	var result [][]string
	for _, data := range courses {
		arrayString, err := data.ArrayString()
		result = append(result, arrayString)
		if err != nil {
			return [][]string{
				{""},
			}, err
		}
	}
	return result, nil
}

func (u Utility) ToFrontCourseData(courseData *course.Course) (*course.FrontCourseData, error) {
	var frontCourseData *course.FrontCourseData
	frontCourseData, err := course.NewFrontCourseData(courseData)
	if err != nil {
		return nil, err
	}
	return frontCourseData, nil
}

// ToCourse Course型に変換
func (u Utility) ToCourse(body *json.Decoder) (*course.Course, error) {
	var frontCourseData course.FrontCourseData
	err := body.Decode(&frontCourseData)
	if err != nil {
		return nil, err
	}
	courseForm, err := course.NewCourse(frontCourseData)
	if err != nil {
		return nil, err
	}
	return courseForm, nil
}
