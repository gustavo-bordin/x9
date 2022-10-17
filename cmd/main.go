package main

import (
	"github.com/gustavo-bordin/x9/x9"
	"github.com/gustavo-bordin/x9/x9/config"
)

func main() {
	cfg := config.GetConfig()
	x9.NewApplication(cfg)
}
