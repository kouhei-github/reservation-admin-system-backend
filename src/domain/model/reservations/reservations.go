package course

import "time"

// Reservations Course DB：reservation.reservations
type Reservations struct {
	StoreId         int       `json:"store_id" gorm:"primary_key"`
	ReserveId       int       `json:"reserve_id" gorm:"primary_key"`
	UserId          int       `json:"user_id"`
	SeatId          int       `json:"seat_id"`
	CourseId        int       `json:"course_id"`
	ReserveNum      int       `json:"reserve_num"`
	ReserveStartDtm time.Time `json:"reserve_start_dtm"`
	ReserveEndDtm   time.Time `json:"reserve_end_dtm"`
	IsHidden        bool      `json:"is_hidden"`
	IsDelete        bool      `json:"is_delete"`
	CreateUser      string    `json:"create_user"`
	CreateAt        time.Time `json:"create_at"`
	UpdateUser      string    `json:"update_user"`
	UpdateAt        time.Time `json:"update_at"`
	Remarks         string    `json:"remarks"`
}
