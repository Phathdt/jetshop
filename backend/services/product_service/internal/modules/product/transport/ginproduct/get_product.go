package ginproduct

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jetshop/common"
	sctx "jetshop/pkg/service-context"
	"jetshop/pkg/service-context/component/gormc"
	"jetshop/pkg/service-context/component/tracing"
	"jetshop/pkg/service-context/core"
	"jetshop/services/product_service/internal/modules/product/business"
	"jetshop/services/product_service/internal/modules/product/repository"
	"jetshop/services/product_service/internal/modules/product/storage"
)

func GetProduct(sc sctx.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracing.StartTrace(ctx, "transport.get")
		defer span.End()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent)

		store := storage.NewSqlStore(db.GetDB())
		repo := repository.NewRepo(store)
		biz := business.NewGetProductBiz(repo)

		book, err := biz.Response(ctx, id)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(book))
	}
}
