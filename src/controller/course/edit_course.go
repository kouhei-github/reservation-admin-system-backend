package course

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/course"
	"net-http/myapp/utils"
	"net/http"
)

type EditCourseHandle struct {
	Service *course.EditCourse
}

func NewEditCourseHandler(s *course.EditCourse) *EditCourseHandle {
	return &EditCourseHandle{Service: s}
}

// EditCourseHandler コース更新コントローラー
func (ru *EditCourseHandle) EditCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}

	jwtToken := r.Header.Get("Authorization")
	// usecase呼び出し
	err := ru.Service.EditCourse(jwtToken, json.NewDecoder(r.Body))
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}
	json.NewEncoder(w).Encode("")
	utils.WriteLogFile("完了しました")
}
