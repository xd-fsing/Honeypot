package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"honeypot/internal/biz"

	pb "honeypot/api/honeypot/v1"
)

type AdminService struct {
	pb.UnimplementedAdminServer

	ac  *biz.AdminUseCase
	log *log.Helper
}

func NewAdminService(ac *biz.AdminUseCase, logger log.Logger) *AdminService {
	return &AdminService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "server/admin")),
	}
}

func (s *AdminService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *AdminService) ListSpy(ctx context.Context, req *pb.ListSpyRequest) (*pb.ListSpyReply, error) {
	return &pb.ListSpyReply{}, nil
}
