package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/ovo/src/api_user/handler"
)

var HandlerSet = wire.NewSet(
	handler.HandlerUserSet,
)
