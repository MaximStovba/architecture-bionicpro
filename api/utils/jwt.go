package utils

import (
	"errors"
	"fmt"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwks         *keyfunc.JWKS
	jwksURL      = "http://localhost:8080/realms/reports-realm/protocol/openid-connect/certs"
	requiredRole = "prothetic_user"
)

// Инициализация при старте приложения
func InitJWKS() error {
	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{})
	return err
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil || !parsedToken.Valid {
		fmt.Println("Could not parse or validate token:", err)
		return errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	if !hasRequiredRole(claims) {
		return errors.New("missing required role")
	}

	return nil
}

func hasRequiredRole(claims jwt.MapClaims) bool {
	realmAccess, ok := claims["realm_access"].(map[string]interface{})
	if !ok {
		return false
	}

	roles, ok := realmAccess["roles"].([]interface{})
	if !ok {
		return false
	}

	for _, role := range roles {
		if roleStr, ok := role.(string); ok && roleStr == requiredRole {
			return true
		}
	}
	return false
}
