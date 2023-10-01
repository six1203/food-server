package service

import (
	pb "food-server/api/food/v1"
	"food-server/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFoodService)

type FoodService struct {
	pb.UnimplementedFoodServer
	uc  *biz.UserUsecase
	cs  *biz.CollectionShopUsecase
	log *log.Helper
}

func NewFoodService(uc *biz.UserUsecase, cs *biz.CollectionShopUsecase, logger log.Logger) *FoodService {
	return &FoodService{uc: uc, cs: cs, log: log.NewHelper(logger)}
}
