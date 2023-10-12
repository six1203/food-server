package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pb "food-server/api/food/v1"
	"food-server/internal/biz"
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
	// 将注入进来的logger实例，用log.NewHelper包装成Helper，绑定到service上，这样就可以在这一层调用这个绑定的的helper对象来打日志了。

	return &FoodService{uc: uc, cs: cs, log: log.NewHelper(logger)}
}
