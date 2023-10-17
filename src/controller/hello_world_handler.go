package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.Header
	fmt.Println(query)
	fmt.Println("TEST")
	fmt.Fprintf(w, "Hello World 1")
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	var data []interface{}
	data = []interface{}{
		"Authテスト", "認証", "認可", 1997, true,
	}

	json.NewEncoder(w).Encode(data)
}
