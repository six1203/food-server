package biz

import (
	"context"
	"fmt"

	v1 "food-server/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
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
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListByHello(context.Context, string) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
	FindByUserinfo(context.Context, string, string) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Username)
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) LoginByUsername(ctx context.Context, username, password string) (*LoginUser, error) {
	user, err := uc.repo.FindByUserinfo(ctx, username, password)
	fmt.Println(username, password, "==================================")
	if err != nil {
		return nil, err
	}
	return &LoginUser{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Token:    "Token",
	}, nil
}
