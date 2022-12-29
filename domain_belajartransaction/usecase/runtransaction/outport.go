package runtransaction

import (
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/repository"
	repository2 "github.com/KuraoHikari/gogen-tx/shared/model/repository"
)

type Outport interface {
	repository2.WithTransactionDB
	repository.SaveProductRepo
	repository.SaveOrderRepo
}
