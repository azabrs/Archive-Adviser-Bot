package telegram

import "Archive-Adviser-Bot/clients/telegram"

type Proccesor struct{
	tg *telegram.Client
	offset int
}