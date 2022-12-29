package entity

import (
	"time"

	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/errorenum"
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/vo"
)

type Order struct {
	ID      vo.OrderID `bson:"_id" json:"id"`
	Created time.Time  `bson:"created" json:"created"`
	User    string		`bson:"user" json:"user"`
}

type OrderCreateRequest struct {
	RandomString string    	`json:"-"`
	Now          time.Time 	`json:"-"`
	User 		string    	`json:"user"`
}

func (r OrderCreateRequest) Validate() error {
	if r.User == "" {
		return errorenum.UserMustNotEmpty
	}
	return nil
}

func NewOrder(req OrderCreateRequest) (*Order, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewOrderID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj Order
	obj.ID = id
	obj.Created = req.Now
	obj.User = req.User

	return &obj, nil
}

type OrderUpdateRequest struct {
	// add field to update here ...
}

func (r *Order) Update(req OrderUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
