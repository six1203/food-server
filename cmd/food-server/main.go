package main

import (
	"flag"
	"fmt"
	"os"

	knacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/protobuf/encoding/protojson"

	"food-server/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "food.server"
	// Version is the version of the compiled software.
	Version = "v1"
	// flagconf is the config flag.
	//flagconf string

	id, _ = os.Hostname()
)

func init() {
	// 设置配置文件路径
	//flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")

	// json序列化配置
	json.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true, //默认值不忽略
		UseProtoNames:   true, //使用proto name返回http字段
	}
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	// 设置nacos配置中心
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "../../configs/nacos/log",
		CacheDir:            "../../configs/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	c := config.New(
		//config.WithSource(
		//file.NewSource(flagconf)),
		config.WithSource(
			knacos.NewConfigSource(
				client,
				knacos.WithGroup("DEFAULT_GROUP"),
				knacos.WithDataID("dev.yaml"),
			),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Config
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if err := c.Watch("userBlacklist", func(key string, value config.Value) {
		fmt.Printf("config changed: %s = %v\n", key, value)
		c.Scan(&bc)
		// 在这里写回调的逻辑
	}); err != nil {
		log.Error(err)
	}

	app, cleanup, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
