package repository

import "net-http/myapp/domain/model/course"

// Course DB：reservation.m_course 操作
type Course struct {
	Name      string `json:"name" binding:"required"`
	CompanyId int    `json:"company_id" binding:"required"`
	CourseId  int    `json:"course_id" binding:"required"`
}

func (c *Course) Select() ([]course.Course, error) {
	var courses []course.Course
	result := db.Table("m_course").Where("company_id = ?", c.CompanyId).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}

func (c *Course) Insert(course course.Course) error {
	result := db.Table("m_course").Create(course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Course) Update(course course.Course) error {
	result := db.Table("m_course").Where("company_id = ? and course_id = ?", c.CompanyId, c.CourseId).Save(&course)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Course) Delete() error {
	result := db.Table("m_course").Where("company_id = ? and course_id = ?", c.CompanyId, c.CourseId).Update("is_delete", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
