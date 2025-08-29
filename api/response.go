package api

import "encoding/json"

type response struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type roomInfo struct {
	RoomID     int `json:"room_id"`
	LiveStatus int `json:"live_status"`
}

type wsConf struct {
	HostServerList []hostServer `json:"host_server_list"`
	Token          string       `json:"token"`
}

type hostServer struct {
	Host string `json:"host"`
}
