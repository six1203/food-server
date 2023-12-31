package data

import (
	slog "log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"food-server/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo, NewCollectionShopRepo)

// Data .
type Data struct {
	db    *gorm.DB
	rdb   *redis.Client
	nacos config_client.IConfigClient
}

// NewData .
func NewData(c *conf.Config, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

func NewDB(c *conf.Config) *gorm.DB {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Data.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, // 表名是否加 s
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	return db
}

func NewRedis(c *conf.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Data.Redis.Addr,
		Password:     c.Data.Redis.Password,
		DB:           int(c.Data.Redis.Db),
		DialTimeout:  c.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Data.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}

//func NewNacos(c *conf.Config) config_client.IConfigClient {
//
//	// 设置nacos配置中心
//	sc := []constant.ServerConfig{
//		*constant.NewServerConfig("127.0.0.1", 8848),
//	}
//
//	cc := &constant.ClientConfig{
//		NamespaceId:         "public",
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "../../configs/nacos/log",
//		CacheDir:            "../../configs/nacos/cache",
//		LogLevel:            "debug",
//	}
//
//	client, err := clients.NewConfigClient(
//		vo.NacosClientParam{
//			ClientConfig:  cc,
//			ServerConfigs: sc,
//		},
//	)
//	if err != nil {
//		panic(err)
//	}
//	return client
//}
