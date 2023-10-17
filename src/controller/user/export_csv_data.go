package user

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/user"
	"net-http/myapp/utils"
	"net/http"
)

type responseBody struct {
	FileName string `json:"file_name" binding:"required"`
}

type ExportCsvUserHandle struct {
	Service *user.UserExportCsvFile
}

func NewUserExportCsvFile(s *user.UserExportCsvFile) *ExportCsvUserHandle {
	return &ExportCsvUserHandle{Service: s}
}

func (exportCsv ExportCsvUserHandle) ExportCsvUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}

	// jwtTokenの認証
	jwtToken := r.Header.Get("Authorization")
	// パラメータやメソッドなどが諸々正しいことが確認できたら、
	// サインイン処理はユースケースに依頼
	fileName, err := exportCsv.Service.ExportCsvFile(jwtToken)
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
	}
	response := responseBody{FileName: fileName}
	json.NewEncoder(w).Encode(response)
	utils.WriteLogFile("完了しました")
}
