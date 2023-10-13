package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopService, NewUserService, NewFoodService)

type FoodService struct {
	Us *UserService
	Ss *ShopService
}

func NewFoodService(us *UserService, s *ShopService) *FoodService {
	// 将注入进来的logger实例，用log.NewHelper包装成Helper，绑定到service上，这样就可以在这一层调用这个绑定的的helper对象来打日志了。
	return &FoodService{Us: us, Ss: s}
}
