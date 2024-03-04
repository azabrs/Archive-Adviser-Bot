package main

import (
	tg_client "Archive-Adviser-Bot/clients/telegram"
	event_consumer "Archive-Adviser-Bot/consumer/event-consumer"
	"Archive-Adviser-Bot/events/telegram"
	"Archive-Adviser-Bot/storage/files"
	"flag"
	"log"
)

var bathSize = 100

func mustToken() string{
	token := flag.String("tg-bot-token", "", "token for acces to tg bot")
	flag.Parse()
	if *token == ""{
		log.Fatal("Token must set")
	}
	return *token
}

func main(){
	tgClient := tg_client.New("api.telegram.org", mustToken())
	eventsProcessor := telegram.New(tgClient, files.New("files_storage"))
	log.Printf("service started")
	cons := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)
	if err := cons.Start(); err != nil{
		log.Fatal("service is stopped")
	}
}