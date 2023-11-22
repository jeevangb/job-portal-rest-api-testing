package cache

import (
	"context"
	"encoding/json"
	"jeevan/jobportal/internal/models"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Rdb struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=service_mocks.go -package=cache
type Caching interface {
	AddToCache(ctx context.Context, jid uint, jdata models.Jobs) error
	GetCahceData(ctx context.Context, jid uint) (string, error)
}

func NewRdbLayer(rdbclnt *redis.Client) Caching {
	return &Rdb{
		rdb: rdbclnt,
	}
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
