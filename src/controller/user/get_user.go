package user

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/user"
	"net-http/myapp/utils"
	"net/http"
)

type GetUserHandle struct {
	Service *user.GetUser
}

func NewGetUserHandler(s *user.GetUser) *GetUserHandle {
	return &GetUserHandle{Service: s}
}

func (ru *GetUserHandle) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}

	jwtToken := r.Header.Get("Authorization")
	// パラメータやメソッドなどが諸々正しいことが確認できたら、
	// サインイン処理はユースケースに依頼
	userData, err := ru.Service.GetUserData(jwtToken)
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}
	userData.Password = ""
	json.NewEncoder(w).Encode(userData)
	utils.WriteLogFile("完了しました")
}
