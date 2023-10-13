package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "food-server/api/food/v1"
	"food-server/internal/biz"
)

type ShopService struct {
	pb.UnimplementedShopServiceServer
	cs  *biz.CollectionShopUsecase
	log *log.Helper
}

func NewShopService(cs *biz.CollectionShopUsecase, logger log.Logger) *ShopService {
	// 将注入进来的logger实例，用log.NewHelper包装成Helper，绑定到service上，这样就可以在这一层调用这个绑定的的helper对象来打日志了。

	return &ShopService{cs: cs, log: log.NewHelper(logger)}
}

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

func (s *ShopService) ListCollectionShop(ctx context.Context, req *pb.ListCollectionShopRequest) (*pb.ListCollectionShopReply, error) {
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

func (s *ShopService) CreateCollectionShop(ctx context.Context, req *pb.CreateCollectionShopRequest) (*pb.CreateCollectionShopReply, error) {
	shop, err := s.cs.CreateCollectionShop(ctx, req.Category, req.Name, req.Logo, req.Address, req.Star)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCollectionShopReply{
		CollectionShop: convertShop(shop),
	}, nil
}

func (s *ShopService) RemoveCollectionShop(ctx context.Context, req *pb.RemoveCollectionShopRequest) (*pb.RemoveCollectionShopReply, error) {
	err := s.cs.RemoveCollectionShop(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveCollectionShopReply{
		Result: "success",
	}, nil
}

func (s *ShopService) GetCollectionShopById(ctx context.Context, req *pb.GetCollectionShopByIdRequest) (*pb.GetCollectionShopByIdReply, error) {
	shop, err := s.cs.GetCollectionShopById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCollectionShopByIdReply{
		CollectionShop: convertShop(shop),
	}, nil
}

func (s *ShopService) UpdateCollectionShopById(ctx context.Context, req *pb.UpdateCollectionShopByIdRequest) (*pb.UpdateCollectionShopByIdReply, error) {
	shop, err := s.cs.UpdateCollectionShopById(ctx, req.Id, req.Star, req.Category, req.Name, req.Logo, req.Address)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCollectionShopByIdReply{
		CollectionShop: convertShop(shop),
	}, nil
}
