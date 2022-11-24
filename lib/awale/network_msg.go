package awale

import "sync"

type messageType int

const (
	connectMessage messageType = iota
	pairMessage
	moveMessage
	errorMessage
)

type networkMsg struct {
	MessageType int `json:"message_type"`

	// connect message
	PinCode int `json:"pin_code"`

	// pair message
	YourTurn int `json:"your_turn"`

	// move message
	SquareIndex int `json:"square_index"`
	Direction   int `json:"direction"`
}

var networkMsgInstance *networkMsg = nil
var networkLock sync.Mutex

func getNetworkMsg() *networkMsg {
	networkLock.Lock()
	var res *networkMsg = nil
	if networkMsgInstance != nil {
		res = &networkMsg{
			networkMsgInstance.MessageType,
			networkMsgInstance.PinCode,
			networkMsgInstance.YourTurn,
			networkMsgInstance.SquareIndex,
			networkMsgInstance.Direction,
		}
	}
	networkLock.Unlock()
	return res
}

func setNetworkMsg(m *networkMsg) {
	networkLock.Lock()
	networkMsgInstance = m
	networkLock.Unlock()
}
