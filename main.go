package main

import (
    "flag"
    "log"
    "net/http"

    "github.com/paxapy/goods/cmd"
)

var mediaPath string

func processFlags() *cmd.Config {
    cfg := &cmd.Config{}

    flag.StringVar(&cfg.ListenSpec, "listen", "localhost:4200", "HTTP listen spec")
    flag.StringVar(
      &cfg.Db.ConnectString,
      "db-connect",
      "host=/var/run/postgresql user=boats dbname=boats sslmode=disable",
      "DB Connect String")
    flag.StringVar(&mediaPath, "media-path", "media", "Path to media dir")

    flag.Parse()
    return cfg
}

func setupHttpMedia(cfg *cmd.Config) {
    log.Printf("Media served from %q.", mediaPath)
    cfg.Api.Media = http.Dir(mediaPath)
}

func main() {
    cfg := processFlags()

    setupHttpMedia(cfg)

    if err := cmd.Run(cfg); err != nil {
        log.Printf("Error in main(): %v", err)
    }
}
