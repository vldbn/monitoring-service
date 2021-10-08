package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

// JWTSecurity JWT token auth
type JWTSecurity struct {
	secretKey string
}

// GenerateToken generates access token.
func (j *JWTSecurity) GenerateToken(username string, exp time.Duration) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(exp).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates JWT token.
func (j *JWTSecurity) ValidateToken(token string) (jwt.MapClaims, error) {
	tokenJWT, err := j.verifyToken(token)
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenJWT.Claims.(jwt.MapClaims); ok && tokenJWT.Valid {
		return claims, nil
	} else {
		e := errors.New("tokenJWT is not valid")
		return nil, e
	}
}

func (j *JWTSecurity) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(j.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken extracts JWT token from request Header
func (j *JWTSecurity) ExtractToken(ctx *fasthttp.RequestCtx) string {
	bearToken := ctx.Request.Header.Peek("Authorization")
	if bearToken == nil {
		return ""
	}
	strArr := strings.Split(string(bearToken), " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// NewJWTSecurity constructor
func NewJWTSecurity(secretKey string) *JWTSecurity {
	return &JWTSecurity{secretKey: secretKey}
}
