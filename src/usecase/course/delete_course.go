package course

import (
	"net-http/myapp/domain"
	"net-http/myapp/utils"
)

type DeleteCourse struct {
	AdminUserRepo    domain.AdminUserRepository
	JwtRepo          domain.AuthJwtToken
	courseRepository domain.CourseRepository
}

// DeleteCourse コース削除ユースケース
func (c DeleteCourse) DeleteCourse(jwtToken string) error {
	userId, err := c.JwtRepo.AuthorizationProcess(jwtToken)
	if err != nil {
		return err
	}

	userData, err := c.AdminUserRepo.FindAdminUserById(userId)
	if err != nil {
		return err
	}

	if userData.IsLogin == false {
		return utils.MyError{Message: "ログインしてください"}
	}

	// DB：delete(論理削除） interfaceを通す
	err = c.courseRepository.Delete()
	if err != nil {
		return err
	}
	return nil
}
