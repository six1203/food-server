package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"food-server/internal/conf"
)

type CollectionShop struct {
	Id        int64
	Category  string
	Name      string
	Logo      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollectionShopRepo interface {
	ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*CollectionShop, int32, error)
	Save(ctx context.Context, category, name, logo, address string) (*CollectionShop, error)
}

type CollectionShopUsecase struct {
	repo   CollectionShopRepo
	log    *log.Helper
	config *conf.Config
}

func NewCollectionShopUsecase(config *conf.Config, repo CollectionShopRepo, logger log.Logger) *CollectionShopUsecase {
	return &CollectionShopUsecase{config: config, repo: repo, log: log.NewHelper(logger)}
}

func (uc *CollectionShopUsecase) ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*CollectionShop, int32, error) {
	return uc.repo.ListCollectionShop(ctx, page, pageSize, fuzzySearchText)
}

func (uc *CollectionShopUsecase) CreateCollectionShop(ctx context.Context, category, name, logo, address string) (*CollectionShop, error) {
	return uc.repo.Save(ctx, category, name, logo, address)
}
