package migrate

import (
	"time"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

const (
	CONCURRENCY_LIMIT = 5
	TIMEOUT           = 20 * time.Second
)

func getMigratePool() (*ants.Pool, error) {
	pool, err := ants.NewPool(CONCURRENCY_LIMIT, ants.WithExpiryDuration(TIMEOUT))
	if err != nil {
		zap.S().Errorw("Error creating migration worker pool", "Error", err)
		return nil, err
	}
	return pool, nil
}
