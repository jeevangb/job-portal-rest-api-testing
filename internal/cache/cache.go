package cache

import (
	"context"
	"encoding/json"
	"errors"
	"jeevan/jobportal/internal/models"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

type Rdb struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=service_mocks.go -package=cache
type Caching interface {
	AddToCache(ctx context.Context, jid uint, jdata models.Jobs) error
	GetCahceData(ctx context.Context, jid uint) (string, error)
	AddToCacheRedis(ctx context.Context, emailKey string, otpValue string) error
	CheckCacheOtp(ctx context.Context, emailKey string) (string, error)
}

func NewRdbLayer(rdbclnt *redis.Client) Caching {
	return &Rdb{
		rdb: rdbclnt,
	}
}

func (c *Rdb) CheckCacheOtp(ctx context.Context, emailKey string) (string, error) {
	otp, err := c.rdb.Get(ctx, emailKey).Result()
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch data from redis cache")
		return "", errors.New("please provide valid email")
	}
	return otp, nil
}
func (c *Rdb) GetCahceData(ctx context.Context, jid uint) (string, error) {

	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := c.rdb.Get(ctx, jobID).Result()
	return str, err

}

func (c *Rdb) AddToCache(ctx context.Context, jid uint, jdata models.Jobs) error {
	jobId := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jdata)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, jobId, val, 0).Err()
	return err

}
func (c *Rdb) AddToCacheRedis(ctx context.Context, emailKey string, otpValue string) error {
	err := c.rdb.Set(ctx, emailKey, otpValue, 3*time.Minute).Err()
	if err != redis.Nil {
		log.Error().Err(err).Msg("failed to set data to the redis cache")
		return errors.New("otp sent failed")
	}
	return redis.Nil

}
