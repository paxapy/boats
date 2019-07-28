package api

import (
    "fmt"
    "log"
    "net"
    "time"
    "net/http"
    "encoding/json"

    "github.com/paxapy/goods/internal/model"
)


type Config struct {
    Assets http.FileSystem
}

func Start(cfg Config, m *model.Model, listener net.Listener) {

    server := &http.Server{
        ReadTimeout:    60 * time.Second,
        WriteTimeout:   60 * time.Second,
        MaxHeaderBytes: 1 << 16}

    http.Handle("/assets/", http.FileServer(cfg.Assets))
    http.Handle("/api/goods/", boatsHandler(m))

    go server.Serve(listener)
}

func listHandler(list interface{}, err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
            log.Print(err)
			http.Error(w, "Error getting list", http.StatusBadRequest)
			return
		}

		js, err := json.Marshal(list)
		if err != nil {
            log.Print(err)
			http.Error(w, "Error encoding json", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, string(js))
	})
}

func boatsHandler(m *model.Model) http.Handler {
    boats, err := m.Boats()
    return listHandler(boats, err)
}

func pagesHandler(m *model.Model) http.Handler {
    pages, err := m.Pages()
    return listHandler(pages, err)
}
