package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/exp/slices"

	"food-server/internal/conf"
	"food-server/internal/pkg/middleware/auth"
)

// User is a User model.
type User struct {
	ID       int64
	Username string
	Password string
}

type LoginUser struct {
	ID       int64
	Username string
	Password string
	Token    string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByUserinfo(context.Context, string, string) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo   UserRepo
	log    *log.Helper
	config *conf.Config
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(config *conf.Config, repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{config: config, repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Username)
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) LoginByUsername(ctx context.Context, username, password string) (*LoginUser, error) {

	fmt.Println("================", uc.config.UserBlacklist)
	if slices.Contains(uc.config.UserBlacklist, username) {
		return nil, errors.New("user is in blacklist")
	}

	user, err := uc.repo.FindByUserinfo(ctx, username, password)
	if err != nil {
		return nil, err
	}
	token := auth.GenerateToken(uc.config.Jwt.Secret, uint(user.ID))
	return &LoginUser{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Token:    token,
	}, nil
}
