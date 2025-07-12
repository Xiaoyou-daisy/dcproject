package pkg

import (
	"fmt"
	//importjwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const APP_KEY = "1342342343"

// TokenHandler是我们获取用户名和密码的处理程序，如果有效，则返回用于将来请求的令牌。
func TokenHandler(userId string) (string, error) {

	// 颁发一个有限期一小时的证书
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(APP_KEY))
	return tokenString, err
}
func GetToken(tokenString string) (jwt.MapClaims, string) {
	//jwt.Parse()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return nil, "登录异常"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("Timing is everything")
			return nil, "登录过期"
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return nil, "登录异常"
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, "登录异常"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, ""
	} else {
		fmt.Println(err)
	}

	return nil, ""
}
