package telegram

import (
	"Archive-Adviser-Bot/lib/e"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)


type Client struct{
	Host string
	Basepath string
	Client http.Client
}

func New(host string, token string) Client{
	return Client{
		Host : host,
		Basepath: newBasePath(token),
		Client: http.Client{},
	}
}



func (c *Client) SendMessage(chat_id int, text string) error{
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chat_id))
	q.Add("text", text)
	_, err := c.doRequest("sendMessage", q)
	if err != nil{
		return e.Wrap("Cant send message", err)
	}
	return nil
}

func (c *Client) Updates(offset int, limit int) (upd []Update, err error){
	defer func(){err = e.WrapIfErr("cant get updates", err)}()
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest("getUpdates", q)
	if err != nil{
		return nil, err
	}
	var res Update_responce
	err = json.Unmarshal(data, &res)
	if err != nil{
		return nil, err
	}
	return res.Result, nil
	
}

func newBasePath(token string) string{
	return "bot" + token
}

func (c *Client) doRequest(method string,q url.Values) (data []byte, err error){
	defer func(){err = e.WrapIfErr("cant do request", err)}()
	u := url.URL{
		Scheme: "http",
		Host: c.Host,
		Path: path.Join(c.Basepath, method),
		RawQuery: q.Encode(),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil{
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil{
		return nil, err
	}
	defer func(){ _ = resp.Body.Close()}()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	return body, nil
}

