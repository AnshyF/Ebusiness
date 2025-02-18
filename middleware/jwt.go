package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secret = []byte("your_jwt_secret_key") // 密钥

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 函数用于解析传入的 JWT 字符串，返回包含用户信息的 Claims 结构体指针和可能出现的错误。
// 参数 tokenString 是待解析的 JWT 字符串。
// 返回值为 *Claims 指针，包含解析后的用户信息，以及可能出现的错误。
func ParseToken(tokenString string) (*Claims, error) {
	// 使用 jwt.ParseWithClaims 函数解析 tokenString，传入自定义的 Claims 结构体和签名密钥的回调函数。
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回预定义的签名密钥，用于验证 JWT 的签名。
		return secret, nil
	})

	// 检查解析后的 token 是否有效，并且其 Claims 是否为 *Claims 类型。
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 如果验证通过，返回解析后的 Claims 结构体指针和 nil 错误。
		return claims, nil
	}
	// 如果验证失败，返回 nil 指针和相应的错误信息。
	return nil, err
}
