package awale

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

// const host string = "178.128.121.220:8080"

const host string = "localhost:3080"
const protocol string = "ws"

var socket *websocket.Conn = nil

func sendMove(move *Move) error {
	return socket.WriteJSON(map[string]interface{}{
		"message_type": moveMessage,
		"square_index": move.Index,
		"direction":    move.Direction,
	})
}

func joinPin(pin int) (int, error) {
	var addr = flag.String("addr", host, "http service address")
	flag.Parse()
	// log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: protocol, Host: *addr, Path: "/pair_client"}
	// log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	socket = c
	if err != nil {
		log.Println("dial:", err)
		return 0, err
	}

	// upon connect
	err = c.WriteJSON(map[string]interface{}{
		"message_type": pairMessage,
		"pin_code":     pin,
	})
	if err != nil {
		log.Println("write:", err)
		c.Close()
		return 0, err
	}

	var body networkMsg
	err = c.ReadJSON(&body)
	if err != nil {
		log.Println("read:", err)
		c.Close()
		return 0, err
	}

	// log.Printf("your turn: %v", body.YourTurn)

	go func() {
		for {
			var body networkMsg
			err := c.ReadJSON(&body)
			if err != nil {
				// log.Println("read:", err)
				c.Close()
				return
			}
			// msg, _ := json.Marshal(body)
			// log.Printf("recv: %s", msg)

			setNetworkMsg(&body)
		}
	}()

	return body.YourTurn, nil
}

func createNewPin() (int, error) {
	var addr = flag.String("addr", host, "http service address")
	flag.Parse()
	// log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: protocol, Host: *addr, Path: "/connect_client"}
	// log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	socket = c
	if err != nil {
		// log.Println("dial:", err)
		return 0, err
	}

	// upon connect
	var body networkMsg
	err = c.ReadJSON(&body)
	if err != nil {
		// log.Println("read:", err)
		c.Close()
		return 0, err
	}

	// log.Printf("code: %v", body.PinCode)

	go func() {
		for {
			var body networkMsg
			err := c.ReadJSON(&body)
			if err != nil {
				// log.Println("read:", err)
				c.Close()
				return
			}

			setNetworkMsg(&body)
			// msg, _ := json.Marshal(body)
			// log.Printf("recv: %s", msg)
		}
	}()

	return body.PinCode, nil
}
