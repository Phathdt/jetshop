package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"jetshop/services/chat_service/internal/modules/chat_transport/chatconsumer"
	"jetshop/shared/appgrpc"
	"jetshop/shared/common"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/discovery/consul"
	"jetshop/shared/sctx/component/ginc"
	smdlw "jetshop/shared/sctx/component/ginc/middleware"
	"jetshop/shared/sctx/component/gormc"
	"jetshop/shared/sctx/component/tracing"
	"jetshop/shared/sctx/component/watermillapp"
	"jetshop/shared/sctx/component/watermillapp/natspub"
	"jetshop/shared/sctx/component/watermillapp/natsrouter"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	serviceName = "chat_detail_thread_consumer"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(consul.NewConsulComponent(common.KeyCompConsul, serviceName, version, 3000)),
		sctx.WithComponent(tracing.NewTracingClient(common.KeyCompJaeger, serviceName, version)),
		sctx.WithComponent(appgrpc.NewChannelClient(common.KeyCompChannelClient)),
		sctx.WithComponent(natsrouter.NewNatsRouter(common.KeyCompNatsSub)),
		sctx.WithComponent(natspub.NewNatsPub(common.KeyCompNatsPub)),
	)
}

var rootCmd = &cobra.Command{
	Use:   serviceName,
	Short: fmt.Sprintf("start %s", serviceName),
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		subscriber := serviceCtx.MustGet(common.KeyCompNatsSub).(watermillapp.Subscriber)
		subscriber.AddNoPublisherHandler("detail_thread", "detail_thread", chatconsumer.PullDetailThreadConsumer(serviceCtx))

		ginComp := serviceCtx.MustGet(common.KeyCompGIN).(ginc.GinComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), gin.Logger(), smdlw.Recovery(serviceCtx), otelgin.Middleware(serviceName))

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

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
