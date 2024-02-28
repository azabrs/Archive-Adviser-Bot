package telegram

type Update_responce struct{
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct{
	ID int `json:"update_id"`
	Message string `json:"message"`
}

