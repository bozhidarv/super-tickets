package auth

import (
	"errors"
	"fmt"
	"strconv"
	"supertickets/internal/models"
	"supertickets/internal/utils"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	Expire int64  `json:"expire"`
}

func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.FormatInt(user.ID, 10),
		"role":    user.Role,
		"expire":  expirationTime.Unix(),
	})
	return token.SignedString([]byte(utils.EnvVars.JwtKey()))
}

func checkTokenAlg(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(utils.EnvVars.JwtKey()), nil
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.Parse(tokenStr, checkTokenAlg)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, err := strconv.ParseInt(claims["user_id"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		role := claims["role"].(string)
		expire := claims["expire"].(float64)
		return &Claims{
			UserID: userId,
			Role:   role,
			Expire: int64(expire),
		}, nil
	} else {
		return nil, errors.New("invalid JWT Claims format")
	}
}
