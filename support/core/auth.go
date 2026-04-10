package core

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/khs1001/gowl-admin/consts"
)

func UserID(ctx context.Context) uint {
	return gconv.Uint(ctx.Value(consts.UserID))
}
