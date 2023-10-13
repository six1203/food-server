package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "food-server/api/food/v1"
	"food-server/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	// 将注入进来的logger实例，用log.NewHelper包装成Helper，绑定到service上，这样就可以在这一层调用这个绑定的的helper对象来打日志了。

	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) LoginByUsername(ctx context.Context, req *pb.LoginByUsernameRequest) (*pb.LoginByUsernameReply, error) {
	user, err := s.uc.LoginByUsername(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginByUsernameReply{
		UserInfo: &pb.UserInfo{Username: user.Username, Password: user.Password, Token: user.Token, Id: user.ID},
	}, nil
}
