package main

import (
    "flag"
    "log"
    "net/http"

    "github.com/paxapy/goods/cmd"
)

var assetsPath string

func processFlags() *cmd.Config {
    cfg := &cmd.Config{}

    flag.StringVar(&cfg.ListenSpec, "listen", "localhost:4200", "HTTP listen spec")
    flag.StringVar(
      &cfg.Db.ConnectString,
      "db-connect",
      "host=/var/run/postgresql user=boats dbname=boats sslmode=disable",
      "DB Connect String")
    flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

    flag.Parse()
    return cfg
}

func setupHttpAssets(cfg *cmd.Config) {
    log.Printf("Assets served from %q.", assetsPath)
    cfg.Api.Assets = http.Dir(assetsPath)
}

func main() {
    cfg := processFlags()

    setupHttpAssets(cfg)

    if err := cmd.Run(cfg); err != nil {
        log.Printf("Error in main(): %v", err)
    }
}
