package user

import (
	"net-http/myapp/domain"
	"net-http/myapp/repository"
	"net-http/myapp/utils"
)

type UserExportCsvFile struct {
	AdminUserRepo domain.AdminUserRepository
	JwtRepo       domain.AuthJwtToken
}

func (export UserExportCsvFile) ExportCsvFile(jwtToken string) (string, error) {
	userId, err := export.JwtRepo.AuthorizationProcess(jwtToken)
	if err != nil {
		return "", err
	}

	userData, err := export.AdminUserRepo.FindAdminUserById(userId)
	if err != nil {
		return "", err
	}

	if userData.IsLogin == false {
		return "", utils.MyError{Message: "ログインしてください"}
	}

	// UserData全て取得
	company := &repository.Company{CompanyId: userData.CompanyId}
	result, err := getUserDataToArrayString(company)
	if err != nil {
		return "", err
	}
	fileName, err := utils.ExportCsv(result)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func getUserDataToArrayString(inter domain.CompanyRepository) ([][]string, error) {
	users, err := inter.GetUserData()
	if err != nil {
		return [][]string{
			{""},
		}, err
	}
	var result [][]string
	for _, user := range users {
		users, err := user.ArrayString()
		result = append(result, users)
		if err != nil {
			return [][]string{
				{""},
			}, err
		}
	}
	return result, nil
}
