package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gusantoniassi/navegante/core/entity/user"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getMockUser() user.User {
	user := user.User{
		ID:       1,
		Name:     "Foobar",
		Email:    "foo@bar.com",
		Password: "123",
	}

	return user
}

func TestJwt_getSecretWithDefinedEnvironmentVariable(t *testing.T) {
	os.Setenv("JWT_SECRET", "foobar")
	secret := getSecret()

	assert.Equal(t, secret, "foobar", "Secret should be the same as the specified environment variable")
}

func TestJwt_getSecretWithUndefinedEnvironmentVariable(t *testing.T) {
	os.Unsetenv("JWT_SECRET")
	secret := getSecret()

	assert.IsType(t, "", secret, "getSecret should return a default string value")
}

func TestGenerateUserToken(t *testing.T) {
	mockUser := getMockUser()

	token, err := GenerateUserToken(mockUser)

	assert.Nilf(t, err, "GenerateUserToken should not return any errors")
	assert.IsType(t, "", token, "GenerateUserToken should return a JWT token")
}

func TestDecodeToken(t *testing.T) {
	mockUser := getMockUser()
	token, err := GenerateUserToken(mockUser)

	assert.Nilf(t, err, "GenerateUserToken should not return any errors")

	usr, err := DecodeUserToken(token)

	assert.Nilf(t, err, "DecodeUserToken should not return any errors")
	assert.IsType(t, &user.User{}, usr, "DecodeUserToken should return a valid user")
	assert.Equal(t, usr.ID, mockUser.ID, "The user ID should match with its value before encoding")
}

func TestDecodeTokenExpired(t *testing.T) {
	// generate an expired token
	claims := jwt.MapClaims{
		"exp":  0,
		"user": getMockUser(),
	}

	token, err := generateToken(claims)
	assert.Nilf(t, err, "generateToken should not return any errors")

	usr, err := DecodeUserToken(token)

	assert.Nilf(t, usr, "User should be null when the token is expired")
	assert.Error(t, err)
}
