package store

import "time"

// Reservations Course DB：reservation.m_store
type store struct {
	StoreId       int       `json:"store_id" gorm:"primary_key"`
	IsParentStore bool      `json:"is_parent_store"`
	ParentStoreId int       `json:"parent_store_id"`
	StoreName     string    `json:"store_name"`
	Address       string    `json:"address"`
	PrefectureId  int       `json:"prefecture_id"`
	City          string    `json:"city"`
	PostCode      int       `json:"postal_code"`
	PhoneNumber   int       `json:"phone_number"`
	Mail          string    `json:"mail"`
	Url           string    `json:"url"`
	IsHidden      bool      `json:"is_hidden"`
	IsDelete      bool      `json:"is_delete"`
	CreateUser    string    `json:"create_user"`
	CreateAt      time.Time `json:"create_at"`
	UpdateUser    string    `json:"update_user"`
	UpdateAt      time.Time `json:"update_at"`
	Remarks       string    `json:"remarks"`
}
