package logic

import (
	"context"

	"gateway-zero-ex/stream/internal/svc"
	"gateway-zero-ex/stream/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StreamLogic {
	return &StreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StreamLogic) Stream(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
