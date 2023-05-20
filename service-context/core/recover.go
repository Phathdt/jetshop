package core

import (
	"jetshop/service-context"
)

func Recover() {
	if r := recover(); r != nil {
		sctx.GlobalLogger().GetLogger("recovered").Errorln(r)
	}
}
