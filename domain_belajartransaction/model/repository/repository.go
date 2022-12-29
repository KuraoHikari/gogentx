package repository

import (
	"context"

	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/entity"
)

type SaveProductRepo interface {
	SaveProduct(ctx context.Context, obj *entity.Product) error
}

type SaveOrderRepo interface {
	SaveOrder(ctx context.Context, obj *entity.Order) error
}
