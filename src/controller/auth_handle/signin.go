package auth_handle

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/auth_usecase"
	"net-http/myapp/utils"
	"net/http"
)

type responseBody struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

type SingninHandle struct {
	Service *auth_usecase.Signin
}

func NewSigninHandler(s *auth_usecase.Signin) *SingninHandle {
	return &SingninHandle{Service: s}
}

func (ru *SingninHandle) SigninHandler(w http.ResponseWriter, r *http.Request) {
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
		Email    string `json:"email" binding:"required"`
		Password string `json:"password"  binding:"required"`
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
	token, err := ru.Service.Signin(input.Email, input.Password)
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}

	response := responseBody{AccessToken: token}
	json.NewEncoder(w).Encode(response)
	utils.WriteLogFile("完了しました")
}
