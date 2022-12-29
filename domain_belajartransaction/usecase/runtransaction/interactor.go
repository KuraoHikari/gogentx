package runtransaction

import (
	"context"

	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/entity"
	"github.com/KuraoHikari/gogen-tx/shared/model/service"
)

type runTransactionInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTransactionInteractor{
		outport: outputPort,
	}
}

func (r *runTransactionInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {
	return service.WithTransaction(ctx, r.outport ,func(ctx context.Context)(*InportResponse, error){
		res := &InportResponse{}
		productObj, err := entity.NewProduct(req.ProductCreateRequest)
		if err != nil {
			return nil, err
		}
		err = r.outport.SaveProduct(ctx, productObj)
		if err != nil {
			return nil, err
		}
		orderObj, err := entity.NewOrder(req.OrderCreateRequest)
		if err != nil {
			return nil, err
		}
		err = r.outport.SaveOrder(ctx, orderObj)
		if err != nil {
			return nil, err
		}
		return res, nil
	})
}
