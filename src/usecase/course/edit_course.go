package course

import (
	"encoding/json"
	"net-http/myapp/domain"
	"net-http/myapp/usecase/course/util"
	"net-http/myapp/utils"
)

type EditCourse struct {
	AdminUserRepo    domain.AdminUserRepository
	JwtRepo          domain.AuthJwtToken
	courseRepository domain.CourseRepository
}

// EditCourse コース更新ユースケース
func (c EditCourse) EditCourse(jwtToken string, parameterBody *json.Decoder) error {
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

	// コース型に変換
	utility := util.NewUtility()
	courseForm, err := utility.ToCourse(parameterBody)
	if err != nil {
		return err
	}

	// DB：update interfaceを通す
	err = c.courseRepository.Update(courseForm)
	if err != nil {
		return err
	}
	return nil
}
