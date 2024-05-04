package sofapp

import (
	"sof/internal/dbops"
	"sof/pkg/db"
	"time"
)


func initDBStore() error {
    opts := &db.MysqlOptions{
        HostName: "localhost",
        Username: "root",
        Password: "123456",
        Database: "sof_db",
        MaxIdleConnections: 10,
        MaxOpenConnections: 5,
        MaxConnectionLifeTime: 10 * time.Second,
    }

    instance, err := db.NewDBConnection(opts)
    if err != nil {
        return err
    }

    _ = dbops.NewDBStore(instance)

    return nil
}

