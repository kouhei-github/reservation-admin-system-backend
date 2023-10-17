package domain

import (
	"net-http/myapp/domain/model/course"
	"net-http/myapp/domain/model/user"
)

// 管理者ユーザーのリポジトリ
type AdminUserRepository interface {
	SaveAdminUser(user *user.AdminUser) error
	FindAdminUserByEmail(email string) (*user.AdminUser, error)
	UpdateAdminUser(user *user.AdminUser) error
	FindAdminUserById(id float64) (*user.AdminUser, error)
}

// JWTTokenの処理
type AuthJwtToken interface {
	CreateJwtToken(serId uint) (string, error)
	AuthorizationProcess(tokenString string) (float64, error)
}

// Companyに紐つく処理
type CompanyRepository interface {
	GetUserData() ([]user.AdminUser, error)
}

// CourseRepository コースに紐づく処理
type CourseRepository interface {
	Select() ([]course.Course, error)
	Insert(course *course.Course) error
	Update(course *course.Course) error
	Delete() error
}
