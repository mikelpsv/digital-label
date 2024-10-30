package main

import (
	"github.com/mikelpsv/digital-label/internal/app"
	"github.com/mikelpsv/digital-label/pkg/config"
)

func main() {
	cfg := config.ReadEnv()
	app.Init(cfg)
}
