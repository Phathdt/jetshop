package redispub

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"jetshop/shared/sctx"
	"jetshop/shared/sctx/component/common"
)

type redisPub struct {
	id        string
	client    *redis.Client
	logger    sctx.Logger
	redisUri  string
	maxActive int
	maxIde    int
}

func NewRedisPub(id string) *redisPub {
	return &redisPub{id: id}
}

func (r *redisPub) ID() string {
	return r.id
}

func (r *redisPub) InitFlags() {
	r.redisUri = common.RedisUri
	r.maxActive = common.MaxActive
	r.maxIde = common.MaxIde
}

func (r *redisPub) Activate(sc sctx.ServiceContext) error {
	r.logger = sctx.GlobalLogger().GetLogger(r.id)
	r.logger.Info("Connecting to Redis at ", r.redisUri, "...")

	opt, err := redis.ParseURL(r.redisUri)

	if err != nil {
		r.logger.Error("Cannot parse Redis ", err.Error())
		return err
	}

	opt.PoolSize = r.maxActive
	opt.MinIdleConns = r.maxIde

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

func (r *redisPub) Stop() error {
	if err := r.client.Close(); err != nil {
		return err
	}

	return nil
}

func (r *redisPub) Publish(ctx context.Context, topic string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = r.client.Publish(ctx, topic, payload).Err(); err != nil {
		return err
	}

	r.logger.Infof("redis pubsub message = %+v\n", string(payload))

	return nil
}
