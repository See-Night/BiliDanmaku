package utils

import "encoding/json"

type WsMsg interface {
	toBytes() []byte
}

type AuthMsg struct {
	UID        int    `json:"uid"`
	RoomID     int    `json:"roomid"`
	Buvid      string `json:"buvid"`
	Protover   int    `json:"protover"`
	SupportAck bool   `json:"support_ack"`
	Scene      string `json:"scene"`
	Platform   string `json:"platform"`
	Type       int    `json:"type"`
	Key        string `json:"key"`
}

func (msg AuthMsg) toBytes() []byte {
	msg_byte, _ := json.Marshal(msg)
	return msg_byte
}

// By default, Bilibili's heartbeat package transmits a string of "[object Object]"
type HeartbeatMsg string

func (msg HeartbeatMsg) toBytes() []byte {
	return []byte(msg)
}

type DanmakuMsg struct {
	Info []any `json:"info"`
}

type ReceiveMsg struct {
	CMD string `json:"cmd"`
}
