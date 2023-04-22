package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义data到biz的返回信息结构
type AdminReply struct {
}

// 与client相关和数据库的交互接口
type AdminRepo interface {
	//保存 管理员用户信息
	SaveAdminInfo(ctx context.Context, name, password string) error
}

type AdminUseCase struct {
	repo AdminRepo
	log  *log.Helper
}

func NewAdminUseCase(repo AdminRepo, logger log.Logger) *AdminUseCase {
	return &AdminUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/admin")),
	}
}
