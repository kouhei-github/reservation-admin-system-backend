package course

import (
	"net-http/myapp/domain"
	"net-http/myapp/domain/model/course"
	"net-http/myapp/usecase/course/util"
	"net-http/myapp/utils"
)

type GetCourse struct {
	AdminUserRepo    domain.AdminUserRepository
	JwtRepo          domain.AuthJwtToken
	courseRepository domain.CourseRepository
}

// GetCourse コース情報取得ユースケース
func (c *GetCourse) GetCourse(jwtToken string) ([]*course.FrontCourseData, error) {
	userId, err := c.JwtRepo.AuthorizationProcess(jwtToken)
	if err != nil {
		return nil, err
	}

	userData, err := c.AdminUserRepo.FindAdminUserById(userId)
	if err != nil {
		return nil, err
	}

	if userData.IsLogin == false {
		return nil, utils.MyError{Message: "ログインしてください"}
	}

	// DB：select interfaceを通す
	courseData, err := c.courseRepository.Select()
	if err != nil {
		return nil, err
	}
	var resultArray []*course.FrontCourseData
	utility := util.NewUtility()
	// TODO FrontCourseData型に変更 上層でエンコードするからここではエンコードしなくていい？
	for _, targetData := range courseData {
		frontCourseData, err := utility.ToFrontCourseData(&targetData)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, frontCourseData)
	}

	return resultArray, nil
}
