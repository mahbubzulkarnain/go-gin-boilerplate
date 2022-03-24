package repository

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// FindByIDs ...
func (r repository) FindByIDs(ctx context.Context, ids []string) (data []*user.Entity, err error) {
	var m *user.Entity
	for _, id := range ids {
		if err = r.redisClient.Get(ctx, r.KeyByID(id)).Scan(m); err != nil && err == redis.Nil {
			err = r.DB.WithContext(ctx).Take(&m).Error
		}
		data = append(data, m)
	}
	if errors.Is(err, redis.Nil) || errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}
