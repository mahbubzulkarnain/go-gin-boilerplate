package repository

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"
)

type repository struct {
	DB          *gorm.DB
	redisClient *redis.Client
}

// NewUserRepository ...
func NewUserRepository(redisClient *redis.Client, db *gorm.DB) user.Repository {
	return &repository{
		DB: db,

		redisClient: redisClient,
	}
}

// KeyByID ...
func (r repository) KeyByID(id string) string {
	return slug.Make(fmt.Sprintf("%s_%s", UserModel{}.TableName(), id))
}
