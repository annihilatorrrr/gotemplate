package main

import (
	"context"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/amarnathcjd/gogram/telegram"
	"github.com/redis/go-redis/v9"
)

var (
	pbot *telegram.Client
	db   *redis.Client
	cotx = context.Background()
)

func handleIfFlood(err error) bool {
	wait := telegram.GetFloodWait(err)
	if wait <= 0 {
		return false
	}
	if wait > 10 {
		return false
	}
	log.Println("flood wait ", wait, "(s), waiting ...")
	time.Sleep(time.Duration(wait+1) * time.Second)
	return true
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	_, _ = msg.Reply(b, "I'm alive!", nil)
	return ext.EndGroups
}

func callrestarter() {
	self, err := os.Executable()
	if err != nil {
		log.Println(err.Error())
		return
	}
	_ = syscall.Exec(self, os.Args, os.Environ())
}

func restart(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	_, _ = msg.Reply(b, "Restarting ...", nil)
	callrestarter()
	return ext.EndGroups
}

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		token = ""
	}
	opt, err := redis.ParseURL("")
	if err != nil {
		log.Fatal(err.Error())
	}
	db = redis.NewClient(opt)
	if err = db.Ping(cotx).Err(); err != nil {
		log.Fatal(err.Error())
	}
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pbot, err = telegram.NewClient(telegram.ClientConfig{AppID: 6, AppHash: "eb06d4abfb49dc3eeb1aeb98ae0f581e", LogLevel: telegram.LogWarn, Session: "pbot.session", FloodHandler: handleIfFlood})
	if err != nil {
		log.Fatal(err)
	}
	pbot.Log.SetColor(false)
	if err = pbot.LoginBot(token); err != nil {
		log.Fatal(err)
	}
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: -1,
	})
	updater := ext.NewUpdater(dispatcher, nil)
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("restart", restart))
	if err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates:    true,
		EnableWebhookDeletion: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			AllowedUpdates: []string{"message"},
		},
	}); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(b.User.FirstName, " has been started!")
	updater.Idle()
	_ = pbot.Stop()
	defer db.Close()
}
