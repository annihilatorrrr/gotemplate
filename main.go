package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		token = "111:3333kkkk"
	}
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			UnhandledErrFunc: func(err error) {
				log.Printf("An error occurred while handling update:\n%s", err.Error())
			},
		}),
	})
	if err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 5,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 5,
			},
		},
	}); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(b.User.FirstName, " has been started!")
	updater.Idle()
}
