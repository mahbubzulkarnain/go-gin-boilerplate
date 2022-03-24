package repository

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

// Create ...
func (r repository) Create(ctx context.Context, e *user.Entity) (rowsAffected int64, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Create(&UserModel{Entity: e})
		if result.Error != nil {
			return result.Error
		}

		var b []byte
		if b, err = json.Marshal(e); err != nil {
			return err
		}

		command := r.redisClient.Set(ctx, r.KeyByID(e.ID), b, 0)
		if err = command.Err(); err != nil {
			return err
		}

		rowsAffected = result.RowsAffected
		return nil
	})
	return
}
