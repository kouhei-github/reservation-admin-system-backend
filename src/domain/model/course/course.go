package course

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Course DB：reservation.m_course
type Course struct {
	StoreId               int       `json:"store_id" gorm:"primary_key"`
	CourseId              int       `json:"course_id" gorm:"primary_key"`
	CourseName            string    `json:"course_name" gorm:"primary_key"`
	CourseKbn             int       `json:"course_kbn"`
	IsCourseNetReserve    bool      `json:"is_course_net_reserve"`
	Price                 int       `json:"price"`
	ItemNum               int       `json:"item_num"`
	Description           string    `json:"description"`
	Contents              string    `json:"contents"`
	FreeDrinkKbn          string    `json:"free_drink_kbn"`
	IsBuffet              bool      `json:"is_buffet"`
	Notes                 string    `json:"notes"`
	Duration              int       `json:"duration"`
	IsAvailablePeople     bool      `json:"is_available_people"`
	AvailableMin          int       `json:"available_min"`
	AvailableMax          int       `json:"available_max"`
	IsAvailableDay        bool      `json:"is_available_day"`
	AvailableDays         string    `json:"available_days"`
	IsAvailableTime       bool      `json:"is_available_time"`
	AvailableTimeMin      time.Time `json:"available_time_min"`
	AvailableTimeMax      time.Time `json:"available_time_max"`
	IsReserveDeadDatetime bool      `json:"is_reserve_dead_datetime"`
	ReserveDeadTimeMin    string    `json:"reserve_dead_time_min"`
	ReserveDeadTimeMax    time.Time `json:"reserve_dead_time_max"`
	IsHidden              bool      `json:"is_hidden"`
	IsDelete              bool      `json:"is_delete"`
}

var timeFormat = "2006-01-02 15:04:05"

func NewCourse(courseFront FrontCourseData) (*Course, error) {
	storeId, err := strconv.Atoi(courseFront.StoreId)
	if err != nil {
		return nil, fmt.Errorf("is not right storeId format")
	}
	var courseId int = 0
	if courseFront.CourseId != "" {
		value, err := strconv.Atoi(courseFront.CourseId)
		if err != nil {
			return nil, fmt.Errorf("is not right courseId format")
		}
		courseId = value
	}
	courseName, err := CheckCharLength(courseFront.CourseName, 50)
	if err != nil {
		return nil, err
	}
	courseKbn, err := NewCourseKbn(courseFront.CourseKbn)
	if err != nil {
		return nil, err
	}
	isCourseNetReserve, err := NewIsCourseNetReserve(courseFront.IsCourseNetReserve)
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(courseFront.Price)
	if err != nil {
		return nil, fmt.Errorf("is not right price format")
	}
	itemNum, err := strconv.Atoi(courseFront.ItemNum)
	if err != nil {
		return nil, fmt.Errorf("is not right itemNum format")
	}
	description, err := CheckCharLength(courseFront.Description, 200)
	if err != nil {
		return nil, err
	}
	contents, err := CheckCharLength(courseFront.Contents, 200)
	if err != nil {
		return nil, err
	}
	freeDrinkKbn, err := NewFreeDrinkKbn(courseFront.IsDrinkCourse, courseFront.IsDrinkCourseOver3h, courseFront.IsDrinkCourseOnly)
	if err != nil {
		return nil, err
	}
	isBuffet, err := strconv.ParseBool(courseFront.IsBuffet)
	if err != nil {
		return nil, err
	}
	notes, err := CheckCharLength(courseFront.Notes, 200)
	if err != nil {
		return nil, err
	}
	duration, err := strconv.Atoi(courseFront.Duration)
	if err != nil {
		return nil, fmt.Errorf("is not right duration format")
	}
	isAvailablePeople, err := strconv.ParseBool(courseFront.IsAvailablePeople)
	if err != nil {
		return nil, err
	}
	availableMin, err := strconv.Atoi(courseFront.AvailableMin)
	if err != nil {
		return nil, fmt.Errorf("is not right availableMin format")
	}
	availableMax, err := strconv.Atoi(courseFront.AvailableMax)
	if err != nil {
		return nil, fmt.Errorf("is not right availableMax format")
	}
	isAvailableDay, err := strconv.ParseBool(courseFront.IsAvailableDay)
	if err != nil {
		return nil, err
	}
	availableDays, err := NewAvailableDays(
		courseFront.AvailableSun,
		courseFront.AvailableMon,
		courseFront.AvailableTue,
		courseFront.AvailableWed,
		courseFront.AvailableThu,
		courseFront.AvailableFri,
		courseFront.AvailableSat,
	)
	if err != nil {
		return nil, err
	}
	isAvailableTime, err := strconv.ParseBool(courseFront.IsAvailableTime)
	if err != nil {
		return nil, err
	}
	availableTimeMin, err := time.Parse(timeFormat, courseFront.AvailableTimeMin)
	if err != nil {
		return nil, err
	}
	availableTimeMax, err := time.Parse(timeFormat, courseFront.AvailableTimeMax)
	if err != nil {
		return nil, err
	}
	isReserveDeadDatetime, err := strconv.ParseBool(courseFront.IsReserveDeadDatetime)
	if err != nil {
		return nil, err
	}
	reserveDeadTimeMin, err := CheckCharLength(courseFront.ReserveDeadTimeMin, 10)
	if err != nil {
		return nil, err
	}
	reserveDeadTimeMax, err := time.Parse(timeFormat, courseFront.ReserveDeadTimeMax)
	if err != nil {
		return nil, err
	}
	isHidden, err := NewIsHidden(courseFront.Status)
	if err != nil {
		return nil, err
	}

	return &Course{
		StoreId:               storeId,
		CourseId:              courseId,
		CourseName:            courseName.value,
		CourseKbn:             courseKbn.value,
		IsCourseNetReserve:    isCourseNetReserve.value,
		Price:                 price,
		ItemNum:               itemNum,
		Description:           description.value,
		Contents:              contents.value,
		FreeDrinkKbn:          freeDrinkKbn.value,
		IsBuffet:              isBuffet,
		Notes:                 notes.value,
		Duration:              duration,
		IsAvailablePeople:     isAvailablePeople,
		AvailableMin:          availableMin,
		AvailableMax:          availableMax,
		IsAvailableDay:        isAvailableDay,
		AvailableDays:         availableDays.value,
		IsAvailableTime:       isAvailableTime,
		AvailableTimeMin:      availableTimeMin,
		AvailableTimeMax:      availableTimeMax,
		IsReserveDeadDatetime: isReserveDeadDatetime,
		ReserveDeadTimeMin:    reserveDeadTimeMin.value,
		ReserveDeadTimeMax:    reserveDeadTimeMax,
		IsHidden:              isHidden.value,
		IsDelete:              false,
	}, nil
}

func (course *Course) ArrayString() ([]string, error) {
	rowData := reflect.ValueOf(course)
	numRow := rowData.NumField()
	var rowDataStr []string
	for i := 0; i < numRow; i++ {
		value := rowData.Field(i).String()
		rowDataStr = append(rowDataStr, value)
	}
	return rowDataStr, nil
}
