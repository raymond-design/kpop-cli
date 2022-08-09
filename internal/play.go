package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var stream beep.StreamSeekCloser

func Play(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http error")
	}

	l_streamer, format, err := mp3.Decode(resp.Body)
	stream = l_streamer
	if err != nil {
		log.Fatal("decoding error")
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(stream)
}

func Stop() {
	stream.Close()
}
