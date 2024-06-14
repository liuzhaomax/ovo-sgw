package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/ovo-sgw/src/api_user/business"
	businessRpc "github.com/liuzhaomax/ovo-sgw/src/api_user_rpc/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessUserSet,
	businessRpc.BusinessUserSet,
)
