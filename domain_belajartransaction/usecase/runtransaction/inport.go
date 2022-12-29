package runtransaction

import (
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/entity"
	"github.com/KuraoHikari/gogen-tx/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.ProductCreateRequest
	entity.OrderCreateRequest
}

type InportResponse struct {
}
