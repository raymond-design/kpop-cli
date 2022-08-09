package types

import "encoding/json"

type SocketRes struct {
	Op int64 `json:"op"`
	D  json.RawMessage
}

type SendData struct {
	Op int64 `json:"op"`
}

type HeartbeatData struct {
	Message   string `json:"message"`
	Heartbeat int64  `json:"heartbeat"`
}

type PlayingData struct {
	Song       Song        `json:"song"`
	Requester  interface{} `json:"requester"`
	Event      interface{} `json:"event"`
	StartTime  string      `json:"startTime"`
	LastPlayed []Song      `json:"lastPlayed"`
	Listeners  int64       `json:"listeners"`
}

type Song struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title"`
	Sources  []interface{} `json:"sources"`
	Artists  []Album       `json:"artists"`
	Albums   []Album       `json:"albums"`
	Duration int64         `json:"duration"`
}

type Album struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	NameRomaji *string `json:"nameRomaji"`
	Image      *string `json:"image"`
}
