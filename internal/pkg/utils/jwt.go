package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = "golang-jwt-secret"

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type TokenInfo struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(JwtSecret),
	}
}

func (j *JWT) GenerateToken(Id uint, Email, Username string) (TokenInfo, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := CustomClaims{
		Id,
		Email,
		Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Vmodev",
		},
	}
	token, err := j.CreateToken(claims)
	return TokenInfo{
		Token:     token,
		ExpiredAt: expireTime,
	}, err

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
