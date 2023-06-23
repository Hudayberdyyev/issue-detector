package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"issue-detector/pkg/app"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "config folder")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Configuration loader
	ctx := context.Background()
	var cfg app.Config
	err := confita.NewLoader(
		file.NewBackend(fmt.Sprintf("%s/default.yaml", configPath)),
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %s\n", err.Error())
		return
	}

	if err := app.Init(cfg); err != nil {
		log.Fatal(err)
	}
}
