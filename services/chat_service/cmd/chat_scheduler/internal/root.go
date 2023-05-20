package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"jetshop/appgrpc"
	"jetshop/common"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/cronjob"
	"jetshop/service-context/component/discovery/consul"
	"jetshop/service-context/component/ginc"
	smdlw "jetshop/service-context/component/ginc/middleware"
	"jetshop/service-context/component/gormc"
	"jetshop/service-context/component/tracing"
	cronjob2 "jetshop/services/chat_service/cmd/chat_scheduler/internal/cronjob"
)

const (
	serviceName = "chat_scheduler"
	version     = "1.0.0"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(ginc.NewGin(common.KeyCompGIN)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(consul.NewConsulComponent(common.KeyCompConsul, serviceName, version, 3000)),
		sctx.WithComponent(tracing.NewTracingClient(common.KeyCompJaeger, serviceName)),
		sctx.WithComponent(cronjob.NewCronjob(common.KeyCron)),
		sctx.WithComponent(appgrpc.NewChannelClient(common.KeyCompChannelClient)),
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

		cronjob2.NewCronjob(serviceCtx)

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
