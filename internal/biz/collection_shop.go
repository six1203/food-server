package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

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
	Star      uint32
}

type CollectionShopRepo interface {
	ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*CollectionShop, int32, error)
	Save(ctx context.Context, star uint32, category, name, logo, address string) (*CollectionShop, error)
	GetFirstByName(ctx context.Context, category, name string) (*CollectionShop, error)
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

func (uc *CollectionShopUsecase) CreateCollectionShop(ctx context.Context, category, name, logo, address string, star uint32) (*CollectionShop, error) {
	shop, err := uc.repo.GetFirstByName(ctx, category, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return uc.repo.Save(ctx, star, category, name, logo, address)
	}
	if shop.Id > 0 {
		return nil, errors.New(fmt.Sprintf("分类：%s下已收藏门店：%s", category, name))
	}
	return nil, err
}
