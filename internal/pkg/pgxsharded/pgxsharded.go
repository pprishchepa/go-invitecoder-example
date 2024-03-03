package pgxsharded

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Cluster struct {
	shards []*pgxpool.Pool
}

func NewCluster(shards []*pgxpool.Pool) *Cluster {
	return &Cluster{
		shards: shards,
	}
}

func (c *Cluster) Size() int {
	return len(c.shards)
}

func (c *Cluster) GetShard(index int) (*pgxpool.Pool, error) {
	if index < 0 || index >= len(c.shards) {
		return nil, errors.New("shard index out of range")
	}

	return c.shards[index], nil
}
