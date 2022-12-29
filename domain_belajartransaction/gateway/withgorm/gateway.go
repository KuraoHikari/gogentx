package withgorm

import (
	"context"
	"fmt"

	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/model/entity"
	"github.com/KuraoHikari/gogen-tx/shared/gogen"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/database"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gateway struct {
	*database.GormWithTransaction
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	//db		*gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "test")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(entity.Product{},entity.Order{})
	return &gateway{
		GormWithTransaction : database.NewGormWithTransaction(db, log),
		log:     log,
		appData: appData,
		config:  cfg,
		//db: db,
	}
}

func (r *gateway) SaveProduct(ctx context.Context, obj *entity.Product) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil{
		return err
	}

	return nil
}

func (r *gateway) SaveOrder(ctx context.Context, obj *entity.Order) error {
	r.log.Info(ctx, "called")

	err := r.ExtractDB(ctx).Save(obj).Error
	if err != nil{
		return err
	}

	return nil
}
