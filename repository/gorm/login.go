package repository

import (
	"context"
	"strings"

	"github.com/mahbubzulkarnain/golang-boilerplate/user"

	"github.com/mahbubzulkarnain/golang-boilerplate/user/dto"
)

// Login ...
func (r repository) Login(ctx context.Context, args dto.LoginRequest) (data *user.Entity, err error) {
	err = r.DB.WithContext(ctx).Select("id, password").Where("LOWER(username)=@user or LOWER(email)=@user", map[string]interface{}{
		"user": strings.ToLower(args.User),
	}).Take(&data).Error
	return
}
