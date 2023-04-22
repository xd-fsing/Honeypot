package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"honeypot/internal/biz"
	"honeypot/internal/data/sqlc/build"
)

var _ biz.ClientRepo = (*clientRepo)(nil)

type clientRepo struct {
	data *Data
	log  *log.Helper
}

func NewClientRepo(data *Data, logger log.Logger) biz.ClientRepo {
	return &clientRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/client")),
	}
}

// SaveLoginSpy /**
func (repo *clientRepo) SaveLoginSpy(ctx context.Context, ip string) error {

	_, err := repo.data.db.SaveLoginSpy(ctx, sqlc.SaveLoginSpyParams{
		Ip:   ip,
		Kind: int16(1),
	})
	if err != nil {
		fmt.Println(err)
	}
	return err
}
