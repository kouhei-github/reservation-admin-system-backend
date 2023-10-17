package repository

import (
	"net-http/myapp/domain/model/user"
)

type Administer struct {
}

func (r *Administer) SaveAdminUser(user *user.AdminUser) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Administer) FindAdminUserByEmail(email string) (*user.AdminUser, error) {
	var adminUser user.AdminUser
	result := db.Find(&adminUser, "email = ?", email)
	if result.Error != nil {
		return &adminUser, result.Error
	}
	return &adminUser, nil
}

func (r *Administer) UpdateAdminUser(user *user.AdminUser) error {
	if result := db.Save(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Administer) FindAdminUserById(id float64) (*user.AdminUser, error) {
	var adminUser user.AdminUser
	if result := db.First(&adminUser, id); result.Error != nil {
		return &adminUser, err
	}
	return &adminUser, nil
}
