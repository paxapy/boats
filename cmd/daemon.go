package cmd

import (
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "github.com/paxapy/goods/api"
    "github.com/paxapy/goods/internal/db"
    "github.com/paxapy/goods/internal/model"
)

type Config struct {
    ListenSpec string
    Db db.Config
    Api api.Config
}

func Run(cfg *Config) error {
    log.Printf("starting, HTTP on: %s\n", cfg.ListenSpec)

    db, err := db.InitDb(cfg.Db)
    if err != nil {
        log.Printf("error initializing database: %v\n", err)
        return err
    }

    dbModel := model.New(db)

    listener, err := net.Listen("tcp", cfg.ListenSpec)
    if err != nil {
        log.Printf("error creating listener: %v\n", err)
    }

    api.Start(cfg.Api, dbModel, listener)
    waitForSignal()

    return nil
}

func waitForSignal() {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    signal := <-ch
    log.Printf("got signal: %v, exiting.", signal)
}
