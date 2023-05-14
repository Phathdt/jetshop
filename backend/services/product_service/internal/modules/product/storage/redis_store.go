package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"jetshop/pkg/service-context/component/tracing"
	"jetshop/pkg/service-context/core"
	"jetshop/services/product_service/internal/modules/product/model"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}

func (s *redisStore) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	ctx, span := tracing.StartTrace(ctx, "redisStore.get")
	defer span.End()

	key := fmt.Sprintf("products/%d", id)

	result, err := s.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, core.ErrRecordNotFound
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	var product model.Product

	err = json.Unmarshal([]byte(result), &product)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &product, nil
}

func (s *redisStore) SetProduct(ctx context.Context, product *model.Product) error {
	ctx, span := tracing.StartTrace(ctx, "redisStore.set")
	defer span.End()

	out, err := json.Marshal(product)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("products/%d", product.Id)
	if err = s.client.Set(ctx, key, out, time.Minute*60).Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
