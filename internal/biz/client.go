package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

// 与client相关和数据库的交互接口
type ClientRepo interface {
	//记录尝试登陆的攻击
	SaveLoginSpy(ctx context.Context, ip string, json *SpyLoginRequestStruct) error
}

type ClientUseCase struct {
	repo ClientRepo
	log  *log.Helper
}

func NewClientUseCase(repo ClientRepo, logger log.Logger) *ClientUseCase {
	return &ClientUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/client")),
	}
}

type SpyLoginRequestStruct struct {
	name     string
	password string
}

// 记录尝试登陆攻击
func (uc *ClientUseCase) SaveSpyLogin(ctx context.Context, name, password string) error {
	var ip string
	if tp, ok := transport.FromServerContext(ctx); ok {
		ip = tp.RequestHeader().Get("x-forwarded-for")
	}

	r := &SpyLoginRequestStruct{}
	r.name = name
	r.password = password

	return uc.repo.SaveLoginSpy(ctx, ip, r)
}
