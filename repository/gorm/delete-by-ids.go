package repository

import (
	"context"
	"gorm.io/gorm"
)

// DeleteByIDs ...
func (r repository) DeleteByIDs(ctx context.Context, ids []string) (rowsAffected int64, err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Delete(&UserModel{}, ids)
		if result.Error != nil {
			return result.Error
		}

		keys := make([]string, 0)
		for _, id := range ids {
			keys = append(keys, r.KeyByID(id))
		}
		command := r.redisClient.Del(ctx, keys...)
		if err = command.Err(); err != nil {
			return err
		}

		rowsAffected = result.RowsAffected
		return nil
	})
	return
}
