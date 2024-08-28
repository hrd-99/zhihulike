package logic

import (
	"context"
	"time"

	"zhihulike/application/user/rpc/internal/code"
	"zhihulike/application/user/rpc/internal/model"
	"zhihulike/application/user/rpc/internal/svc"
	"zhihulike/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	// 当注册名字为空的时候，返回业务自定义错误码
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}
	// 获取插入数据的自增ID
	userId, err := ret.LastInsertId()
	if err != nil {
		// 如果获取自增ID失败，则记录错误日志
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &service.RegisterResponse{UserId: userId}, nil
}
