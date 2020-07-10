// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret         []byte
	defaultExpireTime int64 = 1
	// CookieKey Cookie Key
	CookieKey = "token"
)

// Claims Claims
type Claims struct {
	UserID     int64  `json:"userid"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	RememberMe bool   `json:"rememberMe"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userID int64, username string, expire int64) (string, error) {
	expireTime := genExpireTime(expire)
	claims := Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
func genExpireTime(d int64) time.Time {
	if d == 0 {
		d = defaultExpireTime
	}
	expire := time.Duration(d) * time.Hour
	nowTime := time.Now().Local()
	expireTime := nowTime.Add(expire) // time.Hour
	return expireTime
}
