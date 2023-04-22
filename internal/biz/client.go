package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

// 与client相关和数据库的交互接口
type ClientRepo interface {
	//记录尝试登陆的攻击
	SaveLoginSpy(ctx context.Context, ip string) error
}

type ClientUseCase struct {
	repo ClientRepo
	log  *log.Helper
}

func NewClientUsecase(repo ClientRepo, logger log.Logger) *ClientUseCase {
	return &ClientUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/client")),
	}
}

// 记录尝试登陆攻击
func (cu *ClientUseCase) SaveSpyLogin(ctx context.Context) error {
	var ip string
	if tp, ok := transport.FromServerContext(ctx); ok {
		ip = tp.RequestHeader().Get("x-forwarded-for")
	}
	fmt.Println(ip)
	return cu.repo.SaveLoginSpy(ctx, ip)
}
