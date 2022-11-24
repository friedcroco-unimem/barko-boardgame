package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:3080", "http service address")

type recvMessage struct {
	MessageType int `json:"message_type"`

	// connect message
	PinCode int `json:"pin_code"`

	// pair message
	YourTurn int `json:"your_turn"`

	// move message
	SquareIndex int `json:"square_index"`
	Direction   int `json:"direction"`
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/connect_client"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dial:", err)
		c.Close()
		return
	}

	// upon connect
	var body recvMessage
	err = c.ReadJSON(&body)
	if err != nil {
		log.Println("read:", err)
		c.Close()
		return
	}

	log.Printf("code: %v", body.PinCode)

	done := make(chan recvMessage)

	go func() {
		defer close(done)
		for {
			var body recvMessage
			err := c.ReadJSON(&body)
			if err != nil {
				log.Println("read:", err)
				c.Close()
				return
			}
			msg, _ := json.Marshal(body)
			log.Printf("recv: %s", msg)
		}
	}()

	for {
		select {
		case <-done:
			return
		}
	}
}
