package main

import (
	"bilidanmaku/utils"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var VERSION = "0.0.1 dev"

func printHelp() {
	fmt.Printf("Bili Danmaku %s Copyright © 2024-2025\n", VERSION)
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    bilidanmaku -r <room_id> -u <uid> -s <session_data> -b <buvid3> -c <cookie>")
	fmt.Println("    bilidanmaku -r <room_id> -c <cookie>")
	fmt.Println("    bilidanmaku -h")
	fmt.Println("")
	fmt.Println("Parameters:")
	fmt.Println("    -r <room_id>      Room ID")
	fmt.Println("    -u <uid>          User ID")
	fmt.Println("    -s <session_data> Bilibili seesion data (from cookie)")
	fmt.Println("    -b <buvid3>       Bilibili buvid3 (from cookie)")
	fmt.Println("    -c <cookie>       Bilibili cookies")
	fmt.Println("                      Please enter browser's name or browser's cookies file path")
	fmt.Println("    -h                Print help")
}

func main() {
	// Get command-line parameters
	var room_id = flag.Int("r", 0, "Room ID")
	var uid = flag.Int("u", 0, "User ID")
	var session_data = flag.String("s", "", "Bilibili seesion data (from cookie)")
	var buvid = flag.String("b", "", "Bilibili buvid3 (from cookie)")
	var cookie = flag.String("c", "", "Bilibili cookies, please enter browser's name or browser's cookies file path")
	var help = flag.Bool("h", false, "Print help")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *room_id == 0 {
		log.Println("Please set room id whit `-r`")
		return
	}

	if *cookie != "" {
		// Get cookies from browser
		cookies := make(map[string]string)
		utils.ReadCookies(*cookie, &cookies)

		*uid, _ = strconv.Atoi(cookies["DedeUserID"])
		*session_data = cookies["SESSDATA"]
		*buvid = cookies["buvid3"]
	}

	if *uid == 0 || *session_data == "" || *buvid == "" {
		log.Println("Please set uid, session data and buvid3 whit `-u`, `-s` and `-b`")
		return
	}

	// Get true room ID
	room_true_id, host, key, err := utils.GetWebsocketConf(*room_id, &(map[string]string{
		"SESSDATA": *session_data,
		"buvid3":   *buvid,
	}))
	if err != nil {
		log.Println(err)
		return
	}

	auth_msg := utils.AuthMsg{
		UID:        *uid,
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
