package connect

import (
	"encoding/json"
	"log"
	"time"

	"github.com/raymond-design/kpop-cli/types"
	"github.com/raymond-design/kpop-cli/ui"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var done = false
var ticker *time.Ticker

func sendHeartBeat() {
	data := types.SendData{
		Op: 9,
	}
	conn.WriteJSON(data)
}

func setHeartbeat(repeat int64) {
	sendHeartBeat()
	ticker = time.NewTicker(time.Duration(repeat) * time.Millisecond)
	go func() {
		<-ticker.C
		sendHeartBeat()
	}()
}

func handleMessage(in []byte) {
	var msg types.SocketRes
	json.Unmarshal(in, &msg)
	switch msg.Op {
	case 0:
		var data types.HeartbeatData
		json.Unmarshal(msg.D, &data)
		setHeartbeat(data.Heartbeat)
	case 1:
		var data types.PlayingData
		json.Unmarshal(msg.D, &data)
		album := "None"
		if len(data.Song.Albums) > 0 {
			album = data.Song.Albums[0].Name
		}
		ui.WriteToScreen(data.Song.Title, data.Song.Artists[0].Name, album)
	}
}

func Start(url string) {
	conn_l, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Couldn't connect to websocket")
	}
	conn = conn_l

	go func() {
		for {
			if done {
				conn.Close()
				break
			}
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("Couldn't read websocket message")
			}

			handleMessage(msg)
		}
	}()
}

func Stop() {
	ticker.Stop()
	done = true
}
