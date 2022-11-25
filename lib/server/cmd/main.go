package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var upgrader = websocket.Upgrader{} // use default options

var waitingClients map[int]*websocket.Conn = make(map[int]*websocket.Conn)
var waitingMutex sync.Mutex
var pairedClients map[*websocket.Conn]*websocket.Conn = make(map[*websocket.Conn]*websocket.Conn)
var pairedMutex sync.Mutex

const (
	connectMessage = iota
	pairMessage
	moveMessage
	errorMessage
)

type toMessage struct {
	MessageType int `json:"message_type"`

	// connect message
	PinCode int `json:"pin_code"`

	// pair message
	YourTurn int `json:"your_turn"`

	// move message
	SquareIndex int `json:"square_index"`
	Direction   int `json:"direction"`
}

func StartHTTPServer() error {
	// Create a gin router with default middleware
	path := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Origin", "Cache-Control", "X-Requested-With"}
	path.Use(cors.New(config))

	apiPath := path.Group("")
	registerDataStructureRoute(apiPath)

	return path.Run(":3080")
}

func registerDataStructureRoute(r *gin.RouterGroup) {
	r.GET("/connect_client", connectClient)
	r.GET("/pair_client", pairClient)
}

func connectClient(g *gin.Context) {
	w := g.Writer
	r := g.Request
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// upon connect
	pinCode := int(time.Now().Unix()%10000) + (rand.Intn(90)+10)*10000
	err = c.WriteJSON(map[string]interface{}{
		"message_type": connectMessage,
		"pin_code":     pinCode,
	})
	if err != nil {
		log.Println("write:", err)
		return
	}

	waitingMutex.Lock()
	waitingClients[pinCode] = c
	waitingMutex.Unlock()

	for {
		var body toMessage
		err = c.ReadJSON(&body)
		if err != nil {
			log.Println("marshal:", err)
			return
		}

		// Message type
		if body.MessageType == moveMessage {
			pairedMutex.Lock()
			if val, ok := pairedClients[c]; ok {
				err = val.WriteJSON(map[string]interface{}{
					"message_type": 2,
					"square_index": body.SquareIndex,
					"direction":    body.Direction,
				})
				if err != nil {
					log.Println("write:", err)
				}
			}
			pairedMutex.Unlock()
		}
	}
}

func pairClient(g *gin.Context) {
	w := g.Writer
	r := g.Request

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// upon connect
	var body toMessage
	err = c.ReadJSON(&body)
	if err != nil {
		log.Println("marshal:", err)
		return
	}

	waitingMutex.Lock()
	if val, ok := waitingClients[body.PinCode]; ok {
		delete(waitingClients, body.PinCode)
		waitingMutex.Unlock()

		err = val.WriteJSON(map[string]int{
			"message_type": pairMessage,
			"your_turn":    0,
		})
		if err != nil {
			log.Println("write val:", err)
			return
		}

		pairedMutex.Lock()
		pairedClients[c] = val
		pairedClients[val] = c
		pairedMutex.Unlock()

		err = c.WriteJSON(map[string]int{
			"message_type": pairMessage,
			"your_turn":    1,
		})
		if err != nil {
			log.Println("write c:", err)
			return
		}
	} else {
		waitingMutex.Unlock()
		err = c.WriteJSON(map[string]int{
			"message_type": errorMessage,
			"your_turn":    1,
		})
		if err != nil {
			log.Println("write c:", err)
			return
		}
		return
	}

	for {
		var body toMessage
		err = c.ReadJSON(&body)
		if err != nil {
			log.Println("marshal:", err)
			return
		}

		// Message type
		if body.MessageType == moveMessage {
			pairedMutex.Lock()
			if val, ok := pairedClients[c]; ok {
				err = val.WriteJSON(map[string]interface{}{
					"message_type": 2,
					"square_index": body.SquareIndex,
					"direction":    body.Direction,
				})
				if err != nil {
					log.Println("write:", err)
				}
			}
			pairedMutex.Unlock()
		}
	}
}

func main() {
	// flag.Parse()
	// log.SetFlags(0)
	// http.HandleFunc("/connect_client", connectClient)
	// http.HandleFunc("/pair_client", pairClient)
	// log.Fatal(http.ListenAndServe(*addr, nil))
	StartHTTPServer()
}
