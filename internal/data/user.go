package data

import (
	"context"

	"food-server/internal/biz"

	"food-server/internal/data/mysql/model"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}

func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) FindByUserinfo(ctx context.Context, username, password string) (*biz.User, error) {
	var user model.User
	result := r.data.db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}
