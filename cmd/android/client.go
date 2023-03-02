package main

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"go.uber.org/zap"
	"golang.org/x/net/proxy"
	"os"
	"path/filepath"
)

func Run(f func(ctx context.Context, log *zap.Logger) error) {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func() { _ = log.Sync() }()
	ctx := context.Background()
	if err := f(ctx, log); err != nil {
		log.Fatal("Run failed", zap.Error(err))
	}
}

func main() {
	Run(func(ctx context.Context, log *zap.Logger) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		sessionDir := filepath.Join(home, ".td")
		if err := os.MkdirAll(sessionDir, 0700); err != nil {
			return err
		}
		storage := &telegram.FileSessionStorage{
			Path: filepath.Join(sessionDir, "session-user.json"),
		}

		client := telegram.NewClient(telegram.AndroidAppID, telegram.AndroidAppHash, telegram.Options{
			Logger:         log,
			SessionStorage: storage,
			Resolver:       dcs.Plain(dcs.PlainOptions{Dial: proxy.Dial}),
			DCList:         dcs.Live(),
		})
		return client.Run(ctx, func(ctx context.Context) error {
			flow := auth.NewFlow(
				auth.UsePhone("84927939741"), auth.SendCodeOptions{},
			)
			if err := client.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}
			if _, err := client.Auth().Status(ctx); err != nil {
				return err
			}

			return nil
		})
	})
}
