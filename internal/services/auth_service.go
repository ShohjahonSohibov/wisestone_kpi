package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"kpi/config"
	"kpi/internal/models"
	"kpi/internal/repositories"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	roleRepo *repositories.RoleRepository
}

func NewAuthService(userRepo *repositories.UserRepository, roleRepo *repositories.RoleRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*models.LoginResponse, error) {
    user, err := s.userRepo.FindByUsername(ctx, username)
    if err != nil || user == nil {
        return nil, errors.New("invalid email or password")
    }

    // Compare password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    // Generate JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenStr, err := token.SignedString([]byte(config.Load().Secret))
    if err != nil {
        return nil, errors.New("err in converting token")
    }

    // Add error check for role lookup
    role, err := s.roleRepo.FindByID(ctx, user.RoleId)
		fmt.Println("role:", role, user.RoleId)
    if err != nil {
        return nil, errors.New("error fetching user role")
    }
    if role == nil {
        return nil, errors.New("user role not found")
    }

    roleRes := &models.LoginResponseRole{
        ID:     role.ID,
        NameEn: role.NameEn,
        NameKr: role.NameKr,
    }

    res := &models.LoginResponse{
        Token:  tokenStr,
        RoleId: roleRes,
    }
    return res, nil
}
