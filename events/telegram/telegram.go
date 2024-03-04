package telegram

import (
	"Archive-Adviser-Bot/clients/telegram"
	"Archive-Adviser-Bot/events"
	"Archive-Adviser-Bot/lib/e"
	"Archive-Adviser-Bot/storage"
	"errors"
)

type Meta struct{
	Username string
	ChatID int
}

type Processor struct{
	tg *telegram.Client
	offset int
	storage storage.Storage
}

var (
	ErrUnknownEventType = errors.New("unknow event type")
	ErrUnknownMetaType = errors.New("unknown meta type")	
)

func New(client telegram.Client, storage storage.Storage) *Processor{
	return &Processor{
		tg : &client,
		offset: 0,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int)([]events.Event, error){
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil{
		return nil, e.Wrap("cant get events", err)
	}
	if len(updates) == 0{
		return nil, nil
	}
	res := make([]events.Event, 0, len(updates))

	for _, u := range updates{
		res = append(res, event(u))
	}
	p.offset = updates[len(updates) - 1].ID + 1

	return res, nil
}

func (p *Processor) Procces(event events.Event) error {
	switch event.Type{
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("cant process mesage", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error{
	meta, err := meta(event)
	if err != nil{
		return e.Wrap("cant process message", err)
	}
	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil{
		return  e.Wrap("cant process message", err)
	}

	return nil
}





func event(upd telegram.Update) events.Event{
	updType := fetch_text(upd)
	res := events.Event{
		Text: updType,
		Type: fetch_type(upd),
	}
	if res.Type == events.Message{
		res.Meta = Meta{
			Username: upd.Message.From.Username,
			ChatID: upd.Message.Chat.ID,
		}
	}
	return res
}

func fetch_type(upd telegram.Update) events.Type {
	if upd.Message == nil{
		return events.Unknown
	}
	return events.Message
}

func fetch_text(upd telegram.Update) string {
	if upd.Message == nil{
		return ""
	}
	return upd.Message.Text
}

func meta(event events.Event) (Meta, error){
	res, ok := event.Meta.(Meta)
	if !ok{
		return Meta{}, e.Wrap("cant get Meta", ErrUnknownMetaType)
	}
	return res, nil
}