package auth_usecase

import (
	"net-http/myapp/domain"
	"net-http/myapp/domain/model/user"
	"net-http/myapp/utils"
	"reflect"
)

type Signin struct {
	AdminUserRepo domain.AdminUserRepository
	JwtRepo       domain.AuthJwtToken
}

func (s *Signin) Signin(email string, password string) (string, error) {
	mailVo, err := user.NewEmail(email)
	if err != nil {
		return "", err
	}
	existUser, err := s.AdminUserRepo.FindAdminUserByEmail(mailVo.String())
	if err != nil {
		return "", err
	}
	if reflect.ValueOf(existUser).IsZero() {
		return "", utils.MyError{Message: "メールアドレスが存在しません"}
	}

	// パスワードが一致しているか確認
	passwordVo, err := user.NewPasswrod(password)
	if err != nil {
		return "", err
	}
	isMatch, err := passwordVo.IsMatch(existUser.Password)
	if err != nil {
		return "", err
	}
	if !isMatch {
		return "", utils.MyError{Message: "パスワードが正しくありません"}
	}

	// JWT-Tokenの作成
	newJwtToken, err := s.JwtRepo.CreateJwtToken(existUser.ID)
	if err != nil {
		return "", err
	}

	// ユーザーをログイン中に変更
	existUser.IsLogin = true
	err = s.AdminUserRepo.UpdateAdminUser(existUser)
	if err != nil {
		return "", err
	}

	return newJwtToken, nil
}
