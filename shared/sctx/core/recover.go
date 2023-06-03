package core

import (
	"jetshop/shared/sctx"
)

func Recover() {
	if r := recover(); r != nil {
		sctx.sctx.GlobalLogger().GetLogger("recovered").Errorln(r)
	}
}
