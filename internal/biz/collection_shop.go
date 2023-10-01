package biz

import (
	"context"
	"food-server/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CollectionShop struct {
	Id        int64
	Category  string
	Name      string
	Logo      string
	Address   string
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

type CollectionShopRepo interface {
	ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*CollectionShop, error)
}

type CollectionShopUsecase struct {
	repo   CollectionShopRepo
	log    *log.Helper
	config *conf.Config
}

func NewCollectionShopUsecase(config *conf.Config, repo CollectionShopRepo, logger log.Logger) *CollectionShopUsecase {
	return &CollectionShopUsecase{config: config, repo: repo, log: log.NewHelper(logger)}
}

func (uc *CollectionShopUsecase) ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*CollectionShop, error) {
	return uc.repo.ListCollectionShop(ctx, page, pageSize, fuzzySearchText)
}
