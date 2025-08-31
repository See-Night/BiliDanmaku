package utils

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn

	sendChan    chan []byte
	receiveChan chan []byte

	Done      chan struct{}
	Interrupt chan os.Signal
}

func (c *Client) Init(url string) error {
	// if c.url == "" {
	// 	return errors.New("url is empty")
	// }

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	c.conn = conn

	c.sendChan = make(chan []byte, 10)
	c.receiveChan = make(chan []byte, 100)

	// Set interrupt signal handler
	c.Interrupt = make(chan os.Signal, 1)
	c.Done = make(chan struct{})

	return nil
}

func (c *Client) Start() {
	go c.sendLoop()
	go c.receiveLoop()
}

func (c *Client) sendLoop() {
	for {
		select {
		case msg := <-c.sendChan:
			err := c.conn.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				log.Printf("发送错误: %v", err.Error())
			}
		case <-c.Done:
			return
		}
	}
}

func (c *Client) receiveLoop() {
	for {
		select {
		case <-c.Done:
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("Receive Error: %v", err)
				return
			}
			if message[7] == 3 {
				res, err := brotliDecompress(message[16:])
				if err != nil {
					log.Printf("Brotli Decompress Error: %v", err)
				}

				raws := c.sliceRaws(res)
				for _, raw := range raws {
					c.receiveChan <- raw
				}

			} else {
				c.receiveChan <- message[16:]
			}
		}
	}
}

func (c *Client) Send(msg []byte) {
	c.sendChan <- msg
}

func (c *Client) Receive() []byte {
	return <-c.receiveChan
}

func (c *Client) Close() {
	close(c.Done)
	time.Sleep(100 * time.Millisecond)
	c.conn.Close()
}

func (c *Client) sliceRaws(raw []byte) [][]byte {
	rawSlices := [][]byte{}
	var cursor uint32 = 0
	for uint64(cursor) < uint64(len(raw)) {
		length := bytesToUint32(raw[cursor : cursor+8])
		rawSlices = append(rawSlices, raw[cursor+16:cursor+length])
		cursor += length
	}

	return rawSlices
}
