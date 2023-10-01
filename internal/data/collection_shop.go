package data

import (
	"context"
	"food-server/internal/data/mysql/model"

	"food-server/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (r *CollectionShopRepo) ListCollectionShop(ctx context.Context, page, pageSize int32, fuzzySearchText string) ([]*biz.CollectionShop, error) {
	var shops []*model.CollectionShop
	r.data.db.Where("name like ? or category like ? ", "%"+fuzzySearchText+"%", "%"+fuzzySearchText+"%").Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).Find(&shops)

	var data []*biz.CollectionShop

	for _, shop := range shops {
		data = append(data, &biz.CollectionShop{
			Id:        shop.ID,
			Category:  shop.Category,
			Name:      shop.Name,
			Logo:      shop.Logo,
			Address:   shop.Address,
			CreatedAt: timestamppb.New(shop.CreatedAt),
			UpdatedAt: timestamppb.New(shop.UpdatedAt),
		})
	}
	return data, nil
}
