package services

import (
	"context"
	"sof/pkg/meta"
)

type UserServiceser interface {
    CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error
}

type uservices struct {
}

func (uss *uservices) CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error {
    return nil
}
