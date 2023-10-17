package user

import (
	"fmt"
	"strings"
)

// Email パスワードオブジェクト
type Email struct {
	value string
}

// NewEmail パスワードオブジェクト作成
// ハッシュ化されてない値を扱う
// コンストラクタ
//
// @params pwd パスワード
//
// @return パスワードオブジェクト
func NewEmail(email string) (*Email, error) {
	// emailアドレスのフォーマットが正しいか確認
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("is not right email format")
	}
	return &Email{value: email}, nil
}

// 文字列
// @return
// パスワード
func (pwd *Email) String() string {
	return string(pwd.value)
}
