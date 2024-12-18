package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("Token expired")
	ErrTokenNotValidYet = errors.New("Token is not active yet")
	ErrTokenMalformed   = errors.New("That's not even a token")
	ErrTokenInvalid     = errors.New("Couldn't handle this token")
)

type CustomClaims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

func NewJWT(SigningKey []byte) *JWT {
	return &JWT{
		SigningKey,
	}
}

// CreateToken creates a new token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken parses the token.
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	// verify the token claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
