package service

import (
	"context"
	pb "food-server/api/food/v1"
)

func (s *FoodService) ListCollectionShop(ctx context.Context, req *pb.ListCollectionShopRequest) (*pb.ListCollectionShopReply, error) {
	shopList, err := s.cs.ListCollectionShop(ctx, req.Page, req.PageSize, req.FuzzySearchText)
	if err != nil {
		return nil, err
	}
	var pbShopList []*pb.CollectionShop

	for _, shop := range shopList {
		pbShopList = append(
			pbShopList, &pb.CollectionShop{
				Id:        shop.Id,
				Category:  shop.Category,
				Name:      shop.Name,
				Logo:      shop.Logo,
				Address:   shop.Address,
				CreatedAt: shop.CreatedAt,
				UpdatedAt: shop.UpdatedAt,
			})
	}

	return &pb.ListCollectionShopReply{
		Data:  pbShopList,
		Total: int32(len(pbShopList)),
	}, nil
}
