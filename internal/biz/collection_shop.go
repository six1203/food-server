package biz

import (
	"context"
	"errors"
	"fmt"
	"math"
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
	Create(ctx context.Context, star uint32, category, name, logo, address string) (*CollectionShop, error)
	GetFirstByName(ctx context.Context, category, name string) (*CollectionShop, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*CollectionShop, error)
	Update(ctx context.Context, id int64, star uint32, category, name, logo, address string) (*CollectionShop, error)
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
		return uc.repo.Create(ctx, star, category, name, logo, address)
	}
	if shop.Id > 0 {
		return nil, errors.New(fmt.Sprintf("分类：%s下已收藏门店：%s", category, name))
	}
	return nil, err
}

func (uc *CollectionShopUsecase) RemoveCollectionShop(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *CollectionShopUsecase) GetCollectionShopById(ctx context.Context, id int64) (*CollectionShop, error) {
	return uc.repo.Get(ctx, id)
}

func isBetween1And5(n int) bool {
	return n >= 1 && n <= 5 && float64(n) == math.Floor(float64(n))
}

func (uc *CollectionShopUsecase) UpdateCollectionShopById(ctx context.Context, id int64, star uint32, category, name, logo, address string) (*CollectionShop, error) {
	if !isBetween1And5(int(star)) {
		return nil, errors.New("评分范围为1-5")
	}
	return uc.repo.Update(ctx, id, star, category, name, logo, address)
}
