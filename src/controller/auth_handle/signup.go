package auth_handle

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/auth_usecase"
	"net-http/myapp/utils"
	"net/http"
)

type SignupHandle struct {
	Service *auth_usecase.Signup
}

func NewSignupHandler(s *auth_usecase.Signup) *SignupHandle {
	return &SignupHandle{Service: s}
}

func (ru *SignupHandle) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		// Header情報の追加方法
		//header := w.Header()
		//header.Set("cookie", "aaaa")
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}

	// ユーザのリクエストパラメータを構造体にマッピング
	var input struct {
		Name      string `json:"name" binding:"required"`
		CompanyId int    `json:"company_id" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password"  binding:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response := controller.Response{Status: 401, Text: "入力内容をお確かめください"}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}
	// パラメータやメソッドなどが諸々正しいことが確認できたら、
	// サインイン処理はユースケースに依頼
	err := ru.Service.Signup(input.Name, input.CompanyId, input.Email, input.Password)
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}
	utils.WriteLogFile("完了しました")
}
