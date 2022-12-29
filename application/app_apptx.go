package application

import (
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/controller/resapi"
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/gateway/withgorm"
	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/usecase/runtransaction"
	"github.com/KuraoHikari/gogen-tx/shared/gogen"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/config"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/logger"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/server"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/token"
)

type apptx struct{}

func NewApptx() gogen.Runner {
	return &apptx{}
}

func (apptx) Run() error {

	const appName = "apptx"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := resapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		runtransaction.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
