package service

import (
	"context"
	"errors"

	"github.com/mahbubzulkarnain/golang-boilerplate/user/dto"
)

// Login ...
func (s service) Login(ctx context.Context, args dto.LoginRequest, jwt dto.JWTConfig) (*dto.LoginResponse, error) {
	data, err := s.userRepository.Login(ctx, args)
	if err != nil {
		return nil, err
	}
	if !data.ComparePassword(args.Password) {
		return nil, errors.New("unauthorized")
	}

	var token string
	if token, err = data.Token(jwt.Key, jwt.Duration); err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		UserID: data.ID,
		Token:  token,
	}, nil
}
