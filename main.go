package main

import (
	tgClient "github.com/PavelDonchenko/40projects/simple-telegram-bot/clients/telegram"
	"github.com/PavelDonchenko/40projects/simple-telegram-bot/config"
	"github.com/PavelDonchenko/40projects/simple-telegram-bot/consumer/event-consumer"
	"github.com/PavelDonchenko/40projects/simple-telegram-bot/events/telegram"
	"github.com/PavelDonchenko/40projects/simple-telegram-bot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	cfg := config.MustLoad()
	storage := files.New(storagePath)

	//storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
