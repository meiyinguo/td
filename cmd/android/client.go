package main

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
	"golang.org/x/net/proxy"
	"os"
	"path/filepath"
	"time"
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
			if self, err := client.Self(ctx); err != nil || self.Bot {
				if err := auth.NewFlow(
					auth.UsePhone("8618520533002", 2), auth.SendCodeOptions{
						AllowFlashCall: true,
						CurrentNumber:  true,
						AllowAppHash:   true,
					},
				).Run(ctx, client.Auth()); err != nil {
					return errors.Wrap(err, "auth")
				}
			}

			c := tg.NewClient(client)
			for range time.NewTicker(time.Second * 5).C {
				chats, err := c.MessagesGetAllChats(ctx, nil)

				if d, ok := telegram.AsFloodWait(err); ok {
					// Server told us to wait N seconds before sending next message.
					log.Info("Sleeping", zap.Duration("duration", d))
					time.Sleep(d)
					continue
				}

				if err != nil {
					return errors.Wrap(err, "get chats")
				}

				switch chats.(type) {
				case *tg.MessagesChats: // messages.chats#64ff9fd5
					log.Info("Chats")
				case *tg.MessagesChatsSlice: // messages.chatsSlice#9cd81144
					log.Info("Slice")
				}
			}

			return nil
		})
	})
}
