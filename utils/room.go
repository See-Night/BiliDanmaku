package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

const (
	get_room_id        = "https://api.live.bilibili.com/room/v1/Room/room_init"
	get_websocket_conf = "https://api.live.bilibili.com/room/v1/Danmu/getConf"
)

type HttpResponse struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

func getRoomID(id string) int {
	response, err := get(get_room_id, &(map[string]string{
		"id": id,
	}), nil)
	if err != nil {
		return 0
	}

	var room_info struct {
		RoomID int `json:"room_id"`
	}
	json.Unmarshal(response, &room_info)
	return room_info.RoomID
}

func GetWebsocketConf(id string, cookies *map[string]string) (true_id int, host string, key string, err error) {
	room_id := getRoomID(id)

	response, err := get(get_websocket_conf, &(map[string]string{
		"room_id":  strconv.Itoa(room_id),
		"platform": "pc",
		"player":   "web",
	}), cookies)
	if err != nil {
		return 0, "", "", err
	}

	var ws_conf struct {
		HostServerList []struct {
			Host string `json:"host"`
		} `json:"host_server_list"`
		Token string `json:"token"`
	}
	json.Unmarshal(response, &ws_conf)

	return room_id, ws_conf.HostServerList[0].Host, ws_conf.Token, nil
}

func get(url string, params *map[string]string, cookies *map[string]string) ([]byte, error) {
	// Format URL with parameters
	flag := 1
	if params != nil {
		for key, value := range *params {
			if flag == 1 {
				url += "?" + key + "=" + value
				flag = 0
				continue
			}
			url += "&" + key + "=" + value
		}
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if cookies != nil {
		for key, value := range *cookies {
			request.AddCookie(&http.Cookie{
				Name:  key,
				Value: value,
			})
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp_struct HttpResponse
	json.Unmarshal(body, &resp_struct)

	if resp_struct.Code != 0 {
		return nil, errors.New("Access error")
	}

	return resp_struct.Data, nil
}
