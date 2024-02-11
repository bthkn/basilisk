package main

import (
	"context"

	"github.com/bthkn/basilisk/internal"
	"github.com/bthkn/basilisk/internal/appconfig"
)

func main() {
	cfg, err := appconfig.LoadFromPath(context.Background(), "config/AppConfig.pkl")
	if err != nil {
		panic(err)
	}
	if err = internal.NewServer(cfg).Run(); err != nil {
		panic(err)
	}
}
