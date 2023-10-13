package server

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	food "food-server/api/food/v1"
	"food-server/internal/conf"
	"food-server/internal/pkg/encoder"
	"food-server/internal/pkg/middleware/auth"
	"food-server/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Config, s *service.FoodService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		// 自定义错误响应体
		http.ErrorEncoder(encoder.ErrorEncoder),
		// 自定义统一的返回体
		http.ResponseEncoder(encoder.ResponseEncoder),
		// 中间件
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				auth.JWTAuth(c.Jwt.Secret)).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
		// 浏览器跨域
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)),
	}
	opts = append(opts)
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	food.RegisterUserServiceHTTPServer(srv, s.Us)
	food.RegisterShopServiceHTTPServer(srv, s.Ss)
	return srv
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	// 注意这里的路径不是接口的path，而是接口的完整路径
	whiteList["/api.food.v1.user.UserService/LoginByUsername"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		fmt.Printf("operation================>%v", operation)
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
