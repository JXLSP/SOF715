package dbops

import (
	"database/sql"
	"sof/pkg/db"
	"sync"
)

var (
    DBS *database
    once sync.Once
)

type DBOpser interface {
    DB() *sql.DB
    Users() UserOps
}

type database struct {
    db *sql.DB
}

func NewDBStore(db *sql.DB) *database {
    once.Do(func() {
        DBS = &database{db: db}
    })

    return DBS
}

func (d *database) DB() *sql.DB {
    return db.DBC
}

func (d *database) Users() UserOps {
    return newUserOps(d.db)
}
