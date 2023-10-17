package user

import (
	"gorm.io/gorm"
	"strconv"
)

type AdminUser struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	CompanyId int    `json:"company_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password"  binding:"required"`
	IsLogin   bool   `json:"is_login"`
}

func NewAdminUser(name string, companyId int, email string, password string, isLogin bool) (*AdminUser, error) {
	return &AdminUser{Email: email, Name: name, CompanyId: companyId, Password: password, IsLogin: isLogin}, nil
}

func (usr AdminUser) ArrayString() ([]string, error) {
	userId := strconv.Itoa(int(usr.ID))
	isLogin := "false"
	if usr.IsLogin {
		isLogin = "true"
	}
	return []string{userId, usr.Name, usr.Email, isLogin}, nil
}
