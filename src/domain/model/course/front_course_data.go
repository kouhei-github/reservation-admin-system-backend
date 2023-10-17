package course

import (
	"strconv"
)

// FrontCourseData frontからのデータ
type FrontCourseData struct {
	StoreId               string `json:"store_id"`
	CourseId              string `json:"course_id"`
	Status                string `json:"Status"`
	CourseKbn             string `json:"course_kbn"`
	IsCourseNetReserve    string `json:"is_course_net_reserve"`
	CourseName            string `json:"course_name"`
	Price                 string `json:"price"`
	ItemNum               string `json:"item_num"`
	Description           string `json:"description"`
	Contents              string `json:"contents"`
	IsDrinkCourse         string `json:"is_drink_course"`
	IsDrinkCourseOver3h   string `json:"is_drink_course_over3h"`
	IsDrinkCourseOnly     string `json:"is_drink_course_only"`
	IsBuffet              string `json:"is_buffet"`
	Notes                 string `json:"notes"`
	Duration              string `json:"duration"`
	IsAvailablePeople     string `json:"is_available_people"`
	AvailableMin          string `json:"available_min"`
	AvailableMax          string `json:"available_max"`
	IsAvailableDay        string `json:"is_available_day"`
	AvailableSun          string `json:"available_sun"`
	AvailableMon          string `json:"available_mon"`
	AvailableTue          string `json:"available_tue"`
	AvailableWed          string `json:"available_wed"`
	AvailableThu          string `json:"available_thu"`
	AvailableFri          string `json:"available_fri"`
	AvailableSat          string `json:"available_sat"`
	IsAvailableTime       string `json:"is_available_time"`
	AvailableTimeMin      string `json:"available_time_min"`
	AvailableTimeMax      string `json:"available_time_max"`
	IsReserveDeadDatetime string `json:"is_reserve_dead_datetime"`
	ReserveDeadTimeMin    string `json:"reserve_dead_time_min"`
	ReserveDeadTimeMax    string `json:"reserve_dead_time_max"`
	Expiration            string `json:"expiration"`
}

func NewFrontCourseData(course *Course) (*FrontCourseData, error) {
	storeId := strconv.Itoa(course.StoreId)
	courseId := strconv.Itoa(course.CourseId)
	courseName := course.CourseName
	courseKbn := CourseKbn{value: course.CourseKbn}.ToFrontValue()
	isCourseNetReserve := IsCourseNetReserve{value: course.IsCourseNetReserve}.ToFrontValue()
	price := strconv.Itoa(course.Price)
	itemNum := strconv.Itoa(course.ItemNum)
	description := course.Description
	contents := course.Contents
	freeDrinkKbn := FreeDrinkKbn{value: course.FreeDrinkKbn}.ToFrontValue()
	isDrinkCourse := freeDrinkKbn["isDrinkCourse"]
	isDrinkCourseOver3h := freeDrinkKbn["isDrinkCourseOver3h"]
	isDrinkCourseOnly := freeDrinkKbn["isDrinkCourseOnly"]
	isBuffet := strconv.FormatBool(course.IsBuffet)
	notes := course.Notes
	duration := strconv.Itoa(course.Duration)
	isAvailablePeople := strconv.FormatBool(course.IsAvailablePeople)
	availableMin := strconv.Itoa(course.AvailableMin)
	availableMax := strconv.Itoa(course.AvailableMax)
	isAvailableDay := strconv.FormatBool(course.IsAvailableDay)
	availableDays := AvailableDays{value: course.AvailableDays}.ToFrontValue()
	availableSun := availableDays["availableSun"]
	availableMon := availableDays["availableMon"]
	availableTue := availableDays["availableTue"]
	availableWed := availableDays["availableWed"]
	availableThu := availableDays["availableThu"]
	availableFri := availableDays["availableFri"]
	availableSat := availableDays["availableSat"]
	isAvailableTime := strconv.FormatBool(course.IsAvailableTime)
	availableTimeMin := course.AvailableTimeMin.Format(timeFormat)
	availableTimeMax := course.AvailableTimeMax.Format(timeFormat)
	isReserveDeadDatetime := strconv.FormatBool(course.IsReserveDeadDatetime)
	reserveDeadTimeMin := course.ReserveDeadTimeMin
	reserveDeadTimeMax := course.ReserveDeadTimeMax.Format(timeFormat)
	status := IsHidden{value: course.IsHidden}.ToFrontValue()

	return &FrontCourseData{
		StoreId:               storeId,
		CourseId:              courseId,
		Status:                status,
		CourseKbn:             courseKbn,
		IsCourseNetReserve:    isCourseNetReserve,
		CourseName:            courseName,
		Price:                 price,
		ItemNum:               itemNum,
		Description:           description,
		Contents:              contents,
		IsDrinkCourse:         isDrinkCourse,
		IsDrinkCourseOver3h:   isDrinkCourseOver3h,
		IsDrinkCourseOnly:     isDrinkCourseOnly,
		IsBuffet:              isBuffet,
		Notes:                 notes,
		Duration:              duration,
		IsAvailablePeople:     isAvailablePeople,
		AvailableMin:          availableMin,
		AvailableMax:          availableMax,
		IsAvailableDay:        isAvailableDay,
		AvailableSun:          availableSun,
		AvailableMon:          availableMon,
		AvailableTue:          availableTue,
		AvailableWed:          availableWed,
		AvailableThu:          availableThu,
		AvailableFri:          availableFri,
		AvailableSat:          availableSat,
		IsAvailableTime:       isAvailableTime,
		AvailableTimeMin:      availableTimeMin,
		AvailableTimeMax:      availableTimeMax,
		IsReserveDeadDatetime: isReserveDeadDatetime,
		ReserveDeadTimeMin:    reserveDeadTimeMin,
		ReserveDeadTimeMax:    reserveDeadTimeMax,
		Expiration:            "",
	}, nil
}
