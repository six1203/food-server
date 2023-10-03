package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"food-server/internal/biz"
	"food-server/internal/data/mysql/model"
	pagination "food-server/internal/pkg/util"
)

type CollectionShopRepo struct {
	data *Data
	log  *log.Helper
}

func NewCollectionShopRepo(data *Data, logger log.Logger) biz.CollectionShopRepo {
	return &CollectionShopRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CollectionShopRepo) GetFirstByName(ctx context.Context, category, name string) (*biz.CollectionShop, error) {
	var shop model.CollectionShop
	ret := r.data.db.WithContext(ctx).Where("category = ?", category).Where("name = ?", name).First(&shop)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &biz.CollectionShop{
		Id:        shop.ID,
		Category:  shop.Category,
		Name:      shop.Name,
		Logo:      shop.Logo,
		Address:   shop.Address,
		CreatedAt: shop.CreatedAt,
		UpdatedAt: shop.UpdatedAt,
		Star:      uint32(shop.Star),
	}, nil
}

func (r *CollectionShopRepo) ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*biz.CollectionShop, int32, error) {
	var shops []*model.CollectionShop
	var total int64

	db := r.data.db.WithContext(ctx)

	if fuzzySearchText != "" {
		db = db.Where("name like ? or category like ? ", "%"+fuzzySearchText+"%", "%"+fuzzySearchText+"%")
	}

	// 第一次查询获取总数
	if err := db.Model(&model.CollectionShop{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	result := db.Limit(int(pageSize)).Offset(int(pagination.GetPageOffset(page, pageSize))).Find(&shops)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var data []*biz.CollectionShop

	for _, shop := range shops {
		data = append(data, &biz.CollectionShop{
			Id:        shop.ID,
			Category:  shop.Category,
			Name:      shop.Name,
			Logo:      shop.Logo,
			Address:   shop.Address,
			CreatedAt: shop.CreatedAt,
			UpdatedAt: shop.UpdatedAt,
			Star:      uint32(shop.Star),
		})
	}
	return data, int32(total), nil
}

func (r *CollectionShopRepo) Create(ctx context.Context, star uint32, category, name, logo, address string) (*biz.CollectionShop, error) {
	currentTime := time.Now() // 获取当前时间
	shop := model.CollectionShop{
		Category:  category,
		Name:      name,
		Logo:      logo,
		Address:   address,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Star:      int32(star),
	}
	result := r.data.db.WithContext(ctx).Create(&shop)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.CollectionShop{
		Id:        shop.ID,
		Category:  shop.Category,
		Name:      shop.Name,
		Logo:      shop.Logo,
		Address:   shop.Address,
		CreatedAt: shop.CreatedAt,
		UpdatedAt: shop.UpdatedAt,
		Star:      uint32(shop.Star),
	}, nil
}

func (r *CollectionShopRepo) Delete(ctx context.Context, id int64) error {
	var s model.CollectionShop
	if err := r.data.db.WithContext(ctx).First(&s, id).Error; err != nil {
		return err
	}
	if err := r.data.db.WithContext(ctx).Model(&s).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
