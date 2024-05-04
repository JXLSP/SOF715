package services

import "sof/internal/dbops"

type Serviceser interface {
    Users() UserServiceser
}

type services struct {
    dbs dbops.DBOpser
}

func NewServices(dbs dbops.DBOpser) *services {
    return &services{dbs: dbs}
}

func (sr *services) Users() UserServiceser {
    return NewUserServices(sr.dbs)
}

