package telegram

import (
	"Archive-Adviser-Bot/lib/e"
	"Archive-Adviser-Bot/storage"
	"errors"
	"log"
	"net/url"
	"strings"
)
const (
	Rndcom = "/rnd"
	Helpcom = "/help"
	Startcom = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error{
	text = strings.TrimSpace(text)
	log.Printf("got new command %s from %s", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case Rndcom:
		return p.sendRandom(chatID, username)
	case Helpcom:
		return p.sendHelp(chatID)
	case Startcom:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor)savePage(chatID int, username string, text string) (err error){
	defer func() {err = e.WrapIfErr("cant do command: save page", err)}()

	page := &storage.Page{
		URL : text,
		UserName: username,

	}
	isExist, err := p.storage.IsExist(page)
	if err != nil{
		return err
	}
	if isExist{
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil{
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil{
		return err
	}
	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages){
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages){
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}

	return p.storage.Remove(page)


}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}


func isAddCmd(text string) bool{
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}