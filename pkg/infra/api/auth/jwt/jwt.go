package jwt

import (
	"backend/pkg/identity/accesscontrol"
	"backend/pkg/identity/account"
	"os"
	"time"

	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// Load secret key from environment variable
var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims struct for JWT payload
type Claims struct {
	UserID   int64           `json:"user_id"`
	Username string          `json:"username"`
	UserType account.Account `json:"user_type"`
	// Permissions []accesscontrol.Action `json:"permissions"`
	jwt.RegisteredClaims
}

// TokenProvider interface for JWT operations
type TokenProvider interface {
	GenerateToken(userID int64, username string, userType account.Account, expiry time.Duration) (string, error)
	HasPermission(claims *Claims, requiredPermission accesscontrol.Action) bool
}

// jwtService implements TokenProvider
type jwtService struct{}

// NewTokenProvider initializes the JWT service
func NewTokenProvider() (TokenProvider, error) {
	return &jwtService{}, nil
}

// GenerateToken creates a signed JWT token
func (j *jwtService) GenerateToken(userID int64, username string, userType account.Account, expiry time.Duration) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		UserType: userType,
		// Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth.service",
		},
	}

	// Create and sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken verifies the JWT token and extracts claims
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed with the correct method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and validate claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// HasPermission checks if a user has a specific permission
func (j *jwtService) HasPermission(claims *Claims, requiredPermission accesscontrol.Action) bool {
	// for _, perm := range claims.Permissions {
	// 	if perm == requiredPermission {
	// 		return true
	// 	}
	// }
	return false
}
