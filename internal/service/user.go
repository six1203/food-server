package service

import (
	"context"
	pb "food-server/api/food/v1"
)

func (s *FoodService) LoginByUsername(ctx context.Context, req *pb.LoginByUsernameRequest) (*pb.LoginByUsernameReply, error) {
	user, err := s.uc.LoginByUsername(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginByUsernameReply{
		UserInfo: &pb.UserInfo{Username: user.Username, Password: user.Password, Token: user.Token, Id: user.ID},
	}, nil
}
