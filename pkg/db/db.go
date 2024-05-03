package db

import (
	"database/sql"
	"fmt"
	"time"
    _ "github.com/go-sql-driver/mysql"
)

var (
	DBC *sql.DB
	err error
)

type MysqlOptions struct {
	HostName              string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
}

func (opt *MysqlOptions) GetSDN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=uft8&parseTime=%t&loc=%s`,
		opt.Username,
		opt.Password,
		opt.HostName,
		opt.Database,
		true,
		"Local",
	)
}

func NewDBConnection(opts *MysqlOptions) (*sql.DB, error) {
	DBC, err = sql.Open("mysql", opts.GetSDN())
	if err != nil {
		panic(err.Error())
	}

	DBC.SetMaxIdleConns(opts.MaxIdleConnections)
	DBC.SetMaxOpenConns(opts.MaxOpenConnections)
	DBC.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	return DBC, nil
}

