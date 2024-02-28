package main

import (
	"Archive-Adviser-Bot/clients/telegram"
	"flag"
	"log"
)

func mustToken() string{
	token := flag.String("tg-bot-token", "", "token for acces to tg bot")
	flag.Parse()
	if *token == ""{
		log.Fatal("Token must set")
	}
	return *token
}

func main(){
	tgClient := telegram.New("api.telegram.org", mustToken())
	//fetcher = fetcher.New(tgClient)
	//processor = processor.New(tgClient)
	//consumer(fetcher, processor)
}