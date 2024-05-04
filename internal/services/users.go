package services

import (
	"context"
	"errors"
	"sof/internal/dbops"
	"sof/pkg/meta"
)

type UserServiceser interface {
    CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error
}

type uservices struct {
    dbs dbops.DBOpser
}

func NewUserServices(dbs dbops.DBOpser) *uservices {
    return &uservices{dbs: dbs}
}

func (uss *uservices) CreatedUser(ctx context.Context, request *meta.CreatedUserRequest) error {

    if ok := uss.dbs.Users().ExistUser(ctx, request.UserName); !ok {
        return errors.New("用户已经存在")
    }

    return uss.dbs.Users().CreatedUser(ctx, request)
}

