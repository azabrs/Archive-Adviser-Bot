package telegram

type Update_responce struct{
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct{
	ID int `json:"update_id"`
	Message *incoming_message `json:"message"`
}

type incoming_message struct{
	From User `json:"from"`
	Chat Chat `json:"chat"`
	Text string `json:"text"`
}
 
type User struct{
	Username string `json:"username"`
}

type Chat struct{
	ID int `json:"id"`
}