package gokoreanbots

import "github.com/mailru/easyjson"

//easyjson:skip
// Response is interface of typed response models
type Response interface {
	// GetRaw return raw response
	GetRaw() *RawResponse
}

//easyjson:json
// BotOwner is owner of bot
type BotOwner struct {
	Response

	Id       string   `json:"id"`
	Flags    int      `json:"flags"`
	Github   string   `json:"github"`
	Tag      string   `json:"tag"`
	Username string   `json:"username"`
	Bots     []string `json:"bots"`

	raw *RawResponse
}

func (v BotOwner) GetRaw() *RawResponse {
	return v.raw
}

//easyjson:json
// Bot is bot what registered in Koreanbots
type Bot struct {
	Response

	Id          string     `json:"id"`
	Flags       int        `json:"flags"`
	Owners      []BotOwner `json:"owners"`
	Library     string     `json:"lib"`
	Prefix      string     `json:"prefix"`
	Votes       int        `json:"votes"`
	Servers     int        `json:"servers"`
	Intro       string     `json:"intro"`
	Description string     `json:"desc"`
	Web         string     `json:"web"`
	Git         string     `json:"git"`
	Url         string     `json:"url"`
	Category    []string   `json:"category"`
	Status      string     `json:"status"`
	Discord     string     `json:"discord"`
	State       string     `json:"state"`
	Vanity      string     `json:"vanity"`
	Background  string     `json:"bg"`
	Banner      string     `json:"banner"`
	Tag         string     `json:"tag"`
	Avatar      string     `json:"avatar"`
	Name        string     `json:"name"`

	raw *RawResponse
}

func (v Bot) GetRaw() *RawResponse {
	return v.raw
}

//easyjson:json
// Bots are bots what registered in Koreanbots
type Bots struct {
	Type        string `json:"type"`
	Data        []Bot  `json:"data"`
	CurrentPage int    `json:"currentPage"`
	TotalPage   int    `json:"totalPage"`

	raw *RawResponse
}

func (v Bots) GetRaw() *RawResponse {
	return v.raw
}

//easyjson:json
// Vote is response of (bot/server)'s vote endpoint
type Vote struct {
	Response

	Voted    bool  `json:"voted"`
	LastVote int64 `json:"lastVote"`

	raw *RawResponse
}

func (v Vote) GetRaw() *RawResponse {
	return v.raw
}

//easyjson:json
// RawResponse is raw value of response. You can get this by (Response).GetRaw()
type RawResponse struct {
	Code    int                 `json:"code"`
	Data    easyjson.RawMessage `json:"data"`
	Version int                 `json:"version"`
}

//easyjson:json
// BotStatRequest is body model of bot's stats endpoint
type BotStatRequest struct {
	Servers int `json:"servers"`
	Shards  int `json:"shards"`
}
