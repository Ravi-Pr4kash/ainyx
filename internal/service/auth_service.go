package service

import (
	"context"
	"time"

	"ainyx/db/sqlc"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type AuthService struct {
	db *sqlc.Queries
}

func NewAuthService(db *sqlc.Queries) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (sqlc.AuthUser, error) {
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.AuthUser{}, err
	}

	// Insert user
	user, err := s.db.RegisterUser(ctx, sqlc.RegisterUserParams{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	})

	return user, err
}

func (s *AuthService) GetUserByEmail(ctx context.Context, email string) (sqlc.AuthUser, error) {
	return s.db.GetUserByEmail(ctx, email)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
    user, err := s.db.GetUserByEmail(ctx, email)
    if err != nil {
        return "", err
    }

    // Compare password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", err
    }

    // Create JWT
    secret := os.Getenv("JWT_SECRET")

    claims := jwt.MapClaims{
        "id":    user.ID,
        "name":  user.Name,
        "email": user.Email,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
