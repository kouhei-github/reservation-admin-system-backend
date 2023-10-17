package route

import (
	courseHandling "net-http/myapp/controller/course"
	"net-http/myapp/repository"
	"net-http/myapp/repository/auth_infra"
	"net-http/myapp/usecase/course"
)

func (router *Router) GetCourseRouter() {

	// コース情報取得
	getCourse := courseHandling.NewGetCourseHandler(&course.GetCourse{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/course", getCourse.GetCourseHandler)

	// 新規作成
	createCourse := courseHandling.NewCreateCourseHandler(&course.CreateCourse{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/course/create", createCourse.CreateCourseHandler)

	// 更新
	EditCourse := courseHandling.NewEditCourseHandler(&course.EditCourse{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/course/form", EditCourse.EditCourseHandler)

	// 削除（論理削除）
	DeleteCourse := courseHandling.NewDeleteCourseHandler(&course.DeleteCourse{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/course/delete", DeleteCourse.DeleteCourseHandler)
}
