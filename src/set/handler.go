package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/ovo-sgw/src/api_user/handler"
)

var HandlerSet = wire.NewSet(
	handler.HandlerUserSet,
)
