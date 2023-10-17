package repository

import "net-http/myapp/domain/model/user"

type Company struct {
	Name      string `json:"name" binding:"required"`
	CompanyId int    `json:"company_id" binding:"required"`
}

func (company *Company) GetUserData() ([]user.AdminUser, error) {
	var adminUsers []user.AdminUser
	result := db.Find(&adminUsers, "company_id = ?", company.CompanyId)
	if result.Error != nil {
		return nil, result.Error
	}
	return adminUsers, nil
}
