package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
)

// パスワードオブジェクト
type Password struct {
	value string
}

// NewPasswrod パスワードオブジェクト作成
// ハッシュ化されてない値を扱う
// コンストラクタ
//
// @params pwd パスワード
//
// @return パスワードオブジェクト
func NewPasswrod(pwd string) (*Password, error) {
	// パスワードは５０文字以下であるという仕様がすぐわかる
	if 50 < utf8.RuneCountInString(pwd) {
		return nil, fmt.Errorf("cannot use password over 51 char")
	}
	return &Password{value: pwd}, nil
}

// IsMatch ハッシュ化されたパスワードと一致するか
// @params
// hashPwd ハッシュ化されたパスワード
func (pwd *Password) IsMatch(hashPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd.value))
	if err != nil {
		return false, err
	}
	return true, err
}

// CreateHash ハッシュ化したパスアードを返却
func (pwd *Password) CreateHash() (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(pwd.value), bcrypt.DefaultCost)
	return string(pw), err
}

// 文字列
// @return
// パスワード
func (pwd *Password) String() string {
	return string(pwd.value)
}
