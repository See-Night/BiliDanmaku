package main

import (
	"bilidanmaku/utils"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Get command-line parameters
	var room_id = flag.Int("r", 0, "Room ID")
	var session_data = flag.String("s", "", "Bilibili seesion data (from cookie)")
	var buvid = flag.String("b", "", "Bilibili buvid3 (from cookie)")
	flag.Parse()

	// Get true room ID
	room_true_id, host, key, err := utils.GetWebsocketConf(*room_id, &(map[string]string{
		"SESSDATA": *session_data,
		"buvid3":   *buvid,
	}))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("wss://" + host + "/sub")

	auth_msg := utils.AuthMsg{
		UID:        13390752,
		RoomID:     room_true_id,
		Buvid:      *buvid,
		Protover:   3,
		SupportAck: true,
		Scene:      "room",
		Platform:   "web",
		Key:        key,
		Type:       2,
	}

	auth_pkg := utils.Pkg{}
	auth_pkg.Init(7, auth_msg)

	client := &utils.Client{}
	err = client.Init("wss://" + host + "/sub")
	if err != nil {
		log.Println(err)
		return
	}

	client.Start()
	client.Send(auth_pkg.ToBytes())

	go func() {
		for {
			select {
			case <-client.Done:
				return
			case <-time.After(30 * time.Second):
				sendHeartBeat(client)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-client.Done:
				return
			default:
				msg := client.Receive()
				marshaled, err := marshalMsg(msg)
				if err != nil {
					continue
				}
				log.Printf("%s", marshaled)
			}
		}
	}()

	signal.Notify(client.Interrupt, os.Interrupt)
	<-client.Interrupt

	client.Close()
}

func sendHeartBeat(client *utils.Client) {
	var heartbeat_msg utils.HeartbeatMsg
	heartbeat_msg = "[Object Object]"
	heartbeat_pkg := utils.Pkg{}
	heartbeat_pkg.Init(2, heartbeat_msg)

	client.Send(heartbeat_pkg.ToBytes())
}

func marshalMsg(msg []byte) ([]byte, error) {
	var marshaled []byte
	var err error

	cmd := utils.ReceiveMsg{}
	json.Unmarshal(msg, &cmd)
	if cmd.CMD != "DANMU_MSG" {
		err = errors.New("not a danmu message")
		return nil, err
	}

	danmaku := utils.DanmakuMsg{}
	json.Unmarshal(msg, &danmaku)
	marshaled, err = json.Marshal(danmaku.Info[1])

	return marshaled, err
}
