package redisc

import (
	"context"
	"flag"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	sctx "jetshop/service-context"
)

type RedisComponent interface {
	GetClient() *redis.Client
}

var (
	defaultRedisName      = "DefaultRedis"
	defaultRedisMaxActive = 0 // 0 is unlimited max active connection
	defaultRedisMaxIdle   = 10
)

type RedisDBOpt struct {
	RedisUri  string
	MaxActive int
	MaxIde    int
}

type redisc struct {
	id     string
	client *redis.Client
	logger sctx.Logger
	*RedisDBOpt
}

func NewRedisc(id string) *redisc {
	return &redisc{
		id: id,
		RedisDBOpt: &RedisDBOpt{
			MaxActive: defaultRedisMaxActive,
			MaxIde:    defaultRedisMaxIdle,
		}}
}

func (r *redisc) ID() string {
	return r.id
}

func (r *redisc) InitFlags() {
	flag.StringVar(&r.RedisUri, "redis-uri", "redis://localhost:6379", "(For go-redis) Redis connection-string. Ex: redis://localhost/0")
	flag.IntVar(&r.MaxActive, "redis-pool-max-active", defaultRedisMaxActive, "(For go-redis) Override redis pool MaxActive")
	flag.IntVar(&r.MaxIde, "redis-pool-max-idle", defaultRedisMaxIdle, "(For go-redis) Override redis pool MaxIdle")
}

func (r *redisc) Activate(sc sctx.ServiceContext) error {
	r.logger = sctx.GlobalLogger().GetLogger(r.id)
	r.logger.Info("Connecting to Redis at ", r.RedisUri, "...")

	opt, err := redis.ParseURL(r.RedisUri)

	if err != nil {
		r.logger.Error("Cannot parse Redis ", err.Error())
		return err
	}

	opt.PoolSize = r.MaxActive
	opt.MinIdleConns = r.MaxIde

	client := redis.NewClient(opt)

	// Ping to test Redis connection
	if err = client.Ping(context.Background()).Err(); err != nil {
		r.logger.Error("Cannot connect Redis. ", err.Error())
		return err
	}

	// Enable tracing instrumentation.
	if err = redisotel.InstrumentTracing(client); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err = redisotel.InstrumentMetrics(client); err != nil {
		panic(err)
	}

	// Connect successfully, assign client to goRedisDB
	r.client = client
	return nil
}

func (r *redisc) Stop() error {
	if err := r.client.Close(); err != nil {
		return err
	}

	return nil
}

func (r *redisc) GetClient() *redis.Client {
	return r.client
}
