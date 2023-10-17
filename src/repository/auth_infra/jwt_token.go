package auth_infra

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net-http/myapp/utils"
	"time"
)

type JwtToken struct {
	Exp int64
}

// CreateJwtToken
// JWTトークンの作成
func (j *JwtToken) CreateJwtToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		//"exp":     time.Now().Add(time.Hour * limited).Unix(), // 72時間が有効期限
		"exp": j.Exp,
	}

	// ヘッダーとペイロード生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	if accessToken, err := token.SignedString([]byte("ACCESS_SECRET_KEY")); err != nil {
		return "", err
	} else {
		return accessToken, nil
	}
}

func (j *JwtToken) AuthorizationProcess(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("ACCESS_SECRET_KEY"), nil
	})
	if err != nil {
		return 0, utils.MyError{Message: "tokenの取得に失敗しました"}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, utils.MyError{Message: "認可に失敗しました"}
	}
	j.Exp = int64(claims["exp"].(float64))
	if j.isExpired() {
		return 0, utils.MyError{Message: "トークンの有効期限が切れております"}
	}
	return claims["user_id"].(float64), nil
}

// JWTの期限が切れてるか確認
func (j *JwtToken) isExpired() bool {
	now := time.Now().Unix()
	// 現在時間より期限が小さかったら期限切れ
	if j.Exp < now {
		return true
	}
	return false
}
