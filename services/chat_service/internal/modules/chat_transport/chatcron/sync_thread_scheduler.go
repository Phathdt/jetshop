package chatcron

import (
	"fmt"

	sctx "jetshop/service-context"
)

func SyncThreadScheduler(sc sctx.ServiceContext) func() {
	return func() {
		fmt.Println("hello ba con")
	}
}
