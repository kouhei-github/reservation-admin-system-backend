package auth_usecase

import (
	"net-http/myapp/domain"
	"net-http/myapp/domain/model/user"
	"net-http/myapp/utils"
)

type Signup struct {
	AdminUserRepo domain.AdminUserRepository
}

func (s Signup) Signup(name string, companyId int, email string, password string) error {
	mailVo, err := user.NewEmail(email)
	if err != nil {
		return err
	}
	existUser, err := s.AdminUserRepo.FindAdminUserByEmail(email)

	if err != nil {
		return err
	}
	if existUser.ID != 0 {
		return utils.MyError{Message: "すでにメールアドレスは存在します"}
	}

	passwordVo, err := user.NewPasswrod(password)
	if err != nil {
		return err
	}
	hashPassword, err := passwordVo.CreateHash()
	if err != nil {
		return err
	}
	adminUser, err := user.NewAdminUser(name, companyId, mailVo.String(), hashPassword, false)
	err = s.AdminUserRepo.SaveAdminUser(adminUser)
	if err != nil {
		return err
	}
	return nil
}
