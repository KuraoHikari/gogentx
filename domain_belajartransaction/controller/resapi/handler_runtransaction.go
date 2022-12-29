package resapi

import (
	"context"
	"net/http"
	"time"

	"github.com/KuraoHikari/gogen-tx/domain_belajartransaction/usecase/runtransaction"
	"github.com/KuraoHikari/gogen-tx/shared/gogen"
	"github.com/KuraoHikari/gogen-tx/shared/infrastructure/logger"
	"github.com/KuraoHikari/gogen-tx/shared/model/payload"
	"github.com/KuraoHikari/gogen-tx/shared/util"
	"github.com/gin-gonic/gin"
)

func (r *ginController) runTransactionHandler() gin.HandlerFunc {

	type InportRequest = runtransaction.InportRequest
	type InportResponse = runtransaction.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		// var jsonReq request
		// if err := c.BindJSON(&jsonReq); err != nil {
		// 	r.log.Error(ctx, err.Error())
		// 	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		// 	return
		// }
		id := util.GenerateID(5)
		now := time.Now()

		var req InportRequest

		req.ProductCreateRequest.Now = now
		req.ProductCreateRequest.RandomString = id
		req.OrderCreateRequest.Now = now
		req.OrderCreateRequest.RandomString = id
		req.User = c.DefaultQuery("user", "")


		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
