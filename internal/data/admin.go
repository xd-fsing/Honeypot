package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"honeypot/internal/biz"
)

type adminRepo struct {
	data *Data
	log  *log.Helper
}

func NewAdminRepo(data *Data, logger log.Logger) biz.AdminRepo {
	return &adminRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/role")),
	}
}

/**
 * @Author Fsing
 * @Description 删除角色,删除前需检查是否有用户,否则不能删除
 * @Params
 * @Date 2023/2/6 22:11
 **/
func (repo *adminRepo) SaveAdminInfo(ctx context.Context, name, password string) error {
	//	r, err := repo.data.db.Get(ctx, int(id))

	//	if err != nil {
	//		return nil, err
	//	}

	//	rs, err := r.Update().
	//		SetStatus(false).
	//		Save(ctx)

	return nil

}
