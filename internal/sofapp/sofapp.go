package sofapp

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sof/api/route"
	"sof/pkg/db"
	"syscall"
	"time"
)

func NewAppServe() error {
    return runServe()
}

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

    _ = NewDBStore(instance)

    return nil
}

func runServe() error {
    r := route.InitRoutes()

    httpSrv := startHttpServer(r)

    quit := make(chan os.Signal, 1)

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    if err := httpSrv.Shutdown(ctx); err != nil {
        log.Println("Insecure Server forced to shutdown", err)
        return err
    }

    log.Println("Server exiting")

    return nil
}

func startHttpServer(r http.Handler) *http.Server {
    httpSvr := &http.Server{
        Addr: ":9527",
        Handler: r,
    }

    go func() {
        if err := httpSvr.ListenAndServe(); err != nil {
            log.Println(err.Error())
        }
    }()

    return httpSvr
}
