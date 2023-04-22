package data

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"honeypot/internal/conf"
	"honeypot/internal/data/sqlc/build"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewClientRepo, NewAdminRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *sqlc.Queries
	rdb *redis.Client

	log *log.Helper
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {

	log := log.NewHelper(logger)
	db, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to database: %v", err)
		return nil, nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	queries := sqlc.New(db)

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  queries,
		rdb: rdb,
		log: log,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
