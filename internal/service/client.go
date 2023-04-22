package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "honeypot/api/honeypot/v1"

	"honeypot/internal/biz"
)

type ClientService struct {
	v1.UnimplementedClientServer

	log     *log.Helper
	usecase *biz.ClientUseCase
}

func NewClientService(cc *biz.ClientUseCase, logger log.Logger) *ClientService {
	return &ClientService{
		log:     log.NewHelper(log.With(logger, "module", "server/client")),
		usecase: cc,
	}
}

// 记录用户尝试登陆攻击,将用户的输入转换提交到biz层进行处理,函数名与proto 中定义的名相同
func (cs *ClientService) SpyLogin(ctx context.Context, req *v1.SpyLoginRequest) (reply *v1.SpyLoginReply, err error) {

	reply = &v1.SpyLoginReply{}

	if e := cs.usecase.SaveSpyLogin(ctx); e != nil {
		reply.Code = 1
		reply.Message = e.Error()
	}

	return &v1.SpyLoginReply{
		Code:    0,
		Message: "",
	}, v1.ErrorSaveSpyLoginFailed("保存错误!")

}
