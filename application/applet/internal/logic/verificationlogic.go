package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
	"zhihulike/pkg/util"

	"zhihulike/application/applet/internal/svc"
	"zhihulike/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	prefixVerificationCount = "biz#verification#count#%s" // 用于存储验证码发送次数的Redis键的前缀
	verificationLimitPerDay = 10                          // 每天验证码的发送次数限制
	expireActivation        = 60 * 30                     // 验证码的有效期（30分钟）
)

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	// todo: add your logic here and delete this line

	count, err := l.getVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("getVerificationCount mobile: %s error: %v", req.Mobile, err)
	}
	if count > verificationLimitPerDay {
		return nil, err
	}
	code, err := getActivationCache(req.Mobile, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("getVerificationCount mobile: %s error: %v", req.Mobile, err)
	}
	if len(code) == 0 {
		code = util.RandomNumeric(6)
	}
	//_, err = l.svcCtx.UserRPC.SendSms(l.ctx, &user.SendSmsRequest{
	//	Mobile: req.Mobile,
	//})
	if err != nil {
		logx.Errorf("sendSms mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	err = saveActivationCache(req.Mobile, code, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("saveActivationCache mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	err = l.incrVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("incrVerificationCount mobile: %s error: %v", req.Mobile, err)
	}
	return &types.VerificationResponse{}, nil
}
func (l *VerificationLogic) getVerificationCount(mobile string) (count int, err error) {
	key := fmt.Sprintf(prefixVerificationCount, mobile)
	val, err := l.svcCtx.BizRedis.Get(key)
	if err != nil {
		return 0, err
	}
	if len(val) == 0 {
		return 0, nil
	}
	return strconv.Atoi(val)
}
func getActivationCache(mobile string, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(prefixActivation, mobile)
	return rds.Get(key)
}
func saveActivationCache(mobile string, code string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixActivation, mobile)
	return rds.Setex(key, code, expireActivation)
}

// incrVerificationCount 增加验证次数
func (l *VerificationLogic) incrVerificationCount(mobile string) error {
	key := fmt.Sprintf(prefixVerificationCount, mobile)
	_, err := l.svcCtx.BizRedis.Incr(key)
	if err != nil {
		return err
	}
	return l.svcCtx.BizRedis.Expireat(key, util.EndOfDay(time.Now()).Unix())
}
func delActivationCache(mobile string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixActivation, mobile)
	_, err := rds.Del(key)
	return err
}
