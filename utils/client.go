package utils

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

type ClientCreateMsg struct {
	Conn *websocket.Conn
	Err  error
}

func CreateWsClient(url string) tea.Cmd {
	return func() tea.Msg {
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return ClientCreateMsg{
				Conn: nil,
				Err:  err,
			}
		}
		return ClientCreateMsg{
			Conn: conn,
			Err:  nil,
		}
	}
}

func SliceRaws(raw []byte) [][]byte {
	rawSlices := [][]byte{}
	var cursor uint32 = 0
	for uint64(cursor) < uint64(len(raw)) {
		length := bytesToUint32(raw[cursor : cursor+8])
		rawSlices = append(rawSlices, raw[cursor+16:cursor+length])
		cursor += length
	}

	return rawSlices
}
