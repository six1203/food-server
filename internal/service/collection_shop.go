package service

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "food-server/api/food/v1"
	"food-server/internal/biz"
)

func convertShop(s *biz.CollectionShop) *pb.CollectionShop {
	return &pb.CollectionShop{
		Id:        s.Id,
		Category:  s.Category,
		Name:      s.Name,
		Logo:      s.Logo,
		Address:   s.Address,
		CreatedAt: timestamppb.New(s.CreatedAt),
		UpdatedAt: timestamppb.New(s.UpdatedAt),
		Star:      s.Star,
	}
}

func (s *FoodService) ListCollectionShop(ctx context.Context, req *pb.ListCollectionShopRequest) (*pb.ListCollectionShopReply, error) {
	shopList, total, err := s.cs.ListCollectionShop(ctx, req.Page, req.PageSize, req.FuzzySearchText)

	if err != nil {
		return nil, err
	}
	var pbShopList []*pb.CollectionShop

	for _, shop := range shopList {
		pbShopList = append(pbShopList, convertShop(shop))
	}
	return &pb.ListCollectionShopReply{
		Data:  pbShopList,
		Total: total,
	}, nil
}

func (s *FoodService) CreateCollectionShop(ctx context.Context, req *pb.CreateCollectionShopRequest) (*pb.CreateCollectionShopReply, error) {
	shop, err := s.cs.CreateCollectionShop(ctx, req.Category, req.Name, req.Logo, req.Address, req.Star)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCollectionShopReply{
		CollectionShop: convertShop(shop),
	}, nil
}

func (s *FoodService) RemoveCollectionShop(ctx context.Context, req *pb.RemoveCollectionShopRequest) (*pb.RemoveCollectionShopReply, error) {
	err := s.cs.RemoveCollectionShop(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveCollectionShopReply{
		Result: "success",
	}, nil
}

func (s *FoodService) GetCollectionShopById(ctx context.Context, req *pb.GetCollectionShopByIdRequest) (*pb.GetCollectionShopByIdReply, error) {
	shop, err := s.cs.GetCollectionShopById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCollectionShopByIdReply{
		CollectionShop: convertShop(shop),
	}, nil
}

func (s *FoodService) UpdateCollectionShopById(ctx context.Context, req *pb.UpdateCollectionShopByIdRequest) (*pb.UpdateCollectionShopByIdReply, error) {
	shop, err := s.cs.UpdateCollectionShopById(ctx, req.Id, req.Star, req.Category, req.Name, req.Logo, req.Address)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCollectionShopByIdReply{
		CollectionShop: convertShop(shop),
	}, nil
}
