package sofapp

import (
	"database/sql"
	"sync"
)

var (
    DBS *database
    once sync.Once
)

type database struct {
    db *sql.DB
}

func NewDBStore(db *sql.DB) *database {
    once.Do(func() {
        DBS = &database{db: db}
    })

    return DBS
}

