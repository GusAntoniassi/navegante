package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gusantoniassi/navegante/core/entity/user"
	"log"
	"os"
	"time"
)

const tokenExpiryTime = time.Hour * 8

func getSecret() string {
	var jwtSecret string
	var ok bool

	if jwtSecret, ok = os.LookupEnv("JWT_SECRET"); !ok {
		log.Print("[WARN] No JWT_SECRET environment variable found. You should set it to a random value to improve JWT security")
		jwtSecret = "Navegante is awesome, isn't it? You should really change this though"
	}

	return jwtSecret
}

func generateToken(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(getSecret()))

	return tokenString, err
}

func GenerateUserToken(user user.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = user
	claims["exp"] = time.Now().Add(tokenExpiryTime).Unix()

	token, err := generateToken(claims)
	return token, err
}

func validateTokenExpiry(claims jwt.MapClaims) error {
	exp, ok := claims["exp"].(float64)

	if !ok || int64(exp) < time.Now().Unix() {
		return fmt.Errorf("expired token")
	}

	return nil
}

func getUserFromClaims(claims jwt.MapClaims) (*user.User, error) {
	userMap, ok := claims["user"].(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("no user found in the token")
	}

	id, okId := userMap["ID"].(float64)
	name, okName := userMap["Name"].(string)
	email, okEmail := userMap["Email"].(string)
	password, okPassword := userMap["Password"].(string)

	if !(okId && okName && okEmail && okPassword) {
		return nil, fmt.Errorf("error decoding user values from JWT")
	}

	usr := user.User{
		ID:       user.ID(id),
		Name:     name,
		Email:    email,
		Password: password,
	}

	return &usr, nil
}

func DecodeUserToken(tokenString string) (*user.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(getSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	err = validateTokenExpiry(claims)
	if err != nil {
		return nil, err
	}

	return getUserFromClaims(claims)
}
