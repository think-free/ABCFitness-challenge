package main

import (
	"context"

	"github.com/think-free/ABCFitness-challenge/internal/api"
	"github.com/think-free/ABCFitness-challenge/internal/cliparams"
	"github.com/think-free/ABCFitness-challenge/internal/database"
	"github.com/think-free/ABCFitness-challenge/internal/service"
	"github.com/think-free/ABCFitness-challenge/lib/logging"
)

func main() {
	ctx := context.Background()
	cp := cliparams.New()

	logging.Init(cp.LogLevel)

	db := database.New(ctx, cp)
	srv := service.New(ctx, db)
	ap := api.New(ctx, srv)

	ap.Run()
}
