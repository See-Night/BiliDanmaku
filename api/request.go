package api

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

func GetRoomID(roomID int) (int, error) {
	res, err := get(get_room_id, map[string]string{"id": strconv.Itoa(roomID)}, nil)
	if err != nil {
		return 0, err
	}

	var room_info roomInfo
	json.Unmarshal(res.Data, &room_info)
	if room_info.LiveStatus != 1 {
		return 0, errors.New("直播间关闭")
	}

	return room_info.RoomID, nil
}

func GetWebsocketConf(roomID int, cookies map[string]string) (string, string, error) {
	res, err := get(
		get_websocket_conf,
		map[string]string{
			"room_id":  strconv.Itoa(roomID),
			"platform": "pc",
			"player":   "web",
		},
		&cookies,
	)
	if err != nil {
		return "", "", err
	}

	var conf wsConf
	err = json.Unmarshal(res.Data, &conf)
	if err != nil {
		return "", "", err
	}

	return conf.HostServerList[0].Host, conf.Token, nil
}

func get(url string, params map[string]string, cookies *map[string]string) (*response, error) {
	res := &response{}

	flag := 1
	for key, value := range params {
		if flag == 1 {
			url += "?" + key + "=" + value
			flag = 0
			continue
		}
		url += "&" + key + "=" + value
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if cookies != nil {
		req.Header.Set("Cookie", "SESSDATA="+(*cookies)["SESSDATA"]+"; buvid3="+(*cookies)["buvid3"])
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, res)

	return res, nil
}
