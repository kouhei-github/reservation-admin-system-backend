package user

import (
	"net-http/myapp/domain"
	"net-http/myapp/domain/model/user"
	"net-http/myapp/utils"
)

type GetUser struct {
	AdminUserRepo domain.AdminUserRepository
	JwtRepo       domain.AuthJwtToken
}

func (usr *GetUser) GetUserData(jwtToken string) (*user.AdminUser, error) {
	userId, err := usr.JwtRepo.AuthorizationProcess(jwtToken)
	if err != nil {
		return nil, err
	}

	userData, err := usr.AdminUserRepo.FindAdminUserById(userId)
	if err != nil {
		return nil, err
	}

	if userData.IsLogin == false {
		return nil, utils.MyError{Message: "ログインしてください"}
	}

	return userData, nil
}
