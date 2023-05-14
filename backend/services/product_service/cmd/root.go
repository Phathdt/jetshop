package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"google.golang.org/grpc"
	"jetshop/common"
	sctx "jetshop/pkg/service-context"
	"jetshop/pkg/service-context/component/discovery/consul"
	"jetshop/pkg/service-context/component/ginc"
	smdlw "jetshop/pkg/service-context/component/ginc/middleware"
	"jetshop/pkg/service-context/component/gormc"
	"jetshop/pkg/service-context/component/grpcserverc"
	"jetshop/pkg/service-context/component/migrator"
	"jetshop/pkg/service-context/component/redisc"
	"jetshop/pkg/service-context/component/tracing"
	"jetshop/services/product_service/internal/modules/product/transport/ginproduct"
	"jetshop/services/product_service/internal/modules/product/transport/productgrpc"

	protos "jetshop/proto/out/proto"
)

const (
	serviceName = "product_service"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(redisc.NewRedisc(common.KeyCompRedis)),
		sctx.WithComponent(consul.NewConsulComponent(common.KeyCompConsul, serviceName, version, 3000)),
		sctx.WithComponent(tracing.NewTracingClient(common.KeyCompJaeger, serviceName)),
		sctx.WithComponent(migrator.NewMigrator(common.KeyMigrator)),
		sctx.WithComponent(grpcserverc.NewGrpcServer(common.KeyCompGrpcServer)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: fmt.Sprintf("start %s", serviceName),
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		grpcComp := serviceCtx.MustGet(common.KeyCompGrpcServer).(grpcserverc.GrpcComponent)
		grpcComp.SetRegisterHdl(func(server *grpc.Server) {
			protos.RegisterProductServiceServer(server, productgrpc.NewProductGrpcServer(serviceCtx))
		})

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(ginc.GinComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx), otelgin.Middleware(serviceName))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		apiRouter := router.Group("/api")
		productRouter := apiRouter.Group("/products")
		{
			productRouter.GET("/:id", ginproduct.GetProduct(serviceCtx))
		}

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
