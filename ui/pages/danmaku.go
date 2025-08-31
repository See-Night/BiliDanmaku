package pages

import (
	"bilidanmaku/ui/components"
	"bilidanmaku/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"
)

var screenStyle = lipgloss.NewStyle().Padding(1, 4)

type danmakuPageKeyMap struct {
	Enter key.Binding
	Quit  key.Binding
}

func (k danmakuPageKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Enter,
		k.Quit,
	}
}

func (k danmakuPageKeyMap) FullHelp() [][]key.Binding {
	return nil
}

var danmakuPageKeys = danmakuPageKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("Enter", "submit"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("Esc", "quit"),
	),
}

type errorMsg error
type heartBeatMsg struct {
	Err error
}

func sendHeartBeat(conn *websocket.Conn) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(30 * time.Second)
		var heartbeat_msg utils.HeartbeatMsg
		heartbeat_msg = "[Object Object]"
		heartbeat_pkg := utils.Pkg{}
		heartbeat_pkg.Init(2, heartbeat_msg)

		err := conn.WriteMessage(websocket.BinaryMessage, heartbeat_pkg.ToBytes())

		return heartBeatMsg{
			Err: err,
		}
	}
}

type sendAuthMsg struct {
	Err error
}

func sendAuthPkg(conn *websocket.Conn, authMsg utils.AuthMsg) tea.Cmd {
	return func() tea.Msg {
		auth_pkg := utils.Pkg{}
		auth_pkg.Init(7, authMsg)

		err := conn.WriteMessage(websocket.BinaryMessage, auth_pkg.ToBytes())
		return sendAuthMsg{
			Err: err,
		}
	}
}

type receivedDanmakuMsg struct {
	DanmakuMsg string
	Err        error
}

type receivedDanmakuMsgs struct {
	DanmakuMsgs []receivedDanmakuMsg
}

type sendDanmakuRequest struct {
	Err error
}

func (b *DanmakuPageModel) sendDanmakuMsg(room int, msg []byte) tea.Cmd {
	return func() tea.Msg {
		url := "https://api.live.bilibili.com/msg/send"

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		defer writer.Close()

		fields := map[string]string{
			"bubble":      "0",
			"msg":         string(msg),
			"color":       "16777215",
			"mode":        "1",
			"room_type":   "0",
			"jumpfrom":    "0",
			"reply_mid":   "0",
			"reply_attr":  "0",
			"replay_dmid": "",
			"statistics":  "{\"appId\":100,\"platform\":5}",
			"reply_type":  "0",
			"reply_uname": "",
			"data_extend": "{\"trackid\":\"-99998\"}",
			"fontsize":    "25",
			"rnd":         fmt.Sprintf("%d", time.Now().Unix()),
			"roomid":      strconv.Itoa(room),
			"csrf":        b.bili_jct,
			"csrf_token":  b.bili_jct,
		}

		for key, value := range fields {
			_ = writer.WriteField(key, value)
		}

		writer.Close()
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			return sendDanmakuRequest{
				Err: err,
			}
		}

		req.AddCookie(&http.Cookie{
			Name:  "SESSDATA",
			Value: b.session_data,
		})

		req.Header.Set("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return sendDanmakuRequest{
				Err: err,
			}
		}
		defer resp.Body.Close()

		reqBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return sendDanmakuRequest{
				Err: err,
			}
		}

		var r utils.HttpResponse
		json.Unmarshal(reqBody, &r)

		return sendDanmakuRequest{
			Err: nil,
		}
	}
}

func receiveWsMsg(conn *websocket.Conn) tea.Cmd {
	return func() tea.Msg {
		var danmakus []receivedDanmakuMsg

		_, msg, err := conn.ReadMessage()
		if err != nil {
			danmakus = append(danmakus, receivedDanmakuMsg{
				DanmakuMsg: "",
				Err:        err,
			})
			return receivedDanmakuMsgs{
				DanmakuMsgs: danmakus,
			}
		}

		if msg[7] == 3 {
			res, err := utils.BrotliDecompress(msg[16:])
			if err != nil {
				danmakus = append(danmakus, receivedDanmakuMsg{
					DanmakuMsg: "",
					Err:        err,
				})
				return receivedDanmakuMsgs{
					DanmakuMsgs: danmakus,
				}
			}
			raws := utils.SliceRaws(res)
			for _, raw := range raws {
				danmakus = append(danmakus, receivedDanmakuMsg{
					DanmakuMsg: string(raw),
					Err:        nil,
				})
			}
		} else {
			danmakus = append(danmakus, receivedDanmakuMsg{
				DanmakuMsg: string(msg[16:]),
				Err:        nil,
			})
		}
		return receivedDanmakuMsgs{
			DanmakuMsgs: danmakus,
		}
	}
}

func marshalMsg(msg []byte) ([]byte, []byte, error) {
	var danmakuMsg []byte
	var err error

	cmd := utils.ReceiveMsg{}
	json.Unmarshal(msg, &cmd)
	if cmd.CMD != "DANMU_MSG" {
		err = errors.New("not a danmu message")
		return nil, nil, err
	}

	danmaku := utils.DanmakuMsg{}
	json.Unmarshal(msg, &danmaku)
	danmakuMsg, err = json.Marshal(danmaku.Info[1])
	if err != nil {
		return nil, nil, err
	}

	var user []byte
	user, err = json.Marshal(danmaku.Info[2])
	if err != nil {
		return nil, nil, err
	}

	var userinfo []any
	json.Unmarshal(user, &userinfo)

	var username []byte
	username, err = json.Marshal(userinfo[1])

	return username[1 : len(username)-1], danmakuMsg[1 : len(danmakuMsg)-1], err
}

type DanmakuPageModel struct {
	danmakus []components.DanmakuModel
	inputBar components.InputModel
	help     help.Model

	danmakuslength int

	width  int
	height int

	host         string
	uid          int
	buvid        string
	room         int
	key          string
	bili_jct     string
	session_data string

	conn *websocket.Conn
}

func (d DanmakuPageModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, utils.CreateWsClient("wss://"+d.host+"/sub"))

	for i := range d.danmakus {
		cmds = append(cmds, d.danmakus[i].Init())
	}
	cmds = append(cmds, d.inputBar.Init())
	return tea.Batch(cmds...)
}

func (d DanmakuPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		d.width = msg.Width
		d.height = msg.Height
		d.help.Width = msg.Width - 8
		if d.danmakuslength > msg.Height-5 {
			d.adaptiveDanmakuLength(d.danmakuslength - msg.Height + 7)
		}
		d.danmakuslength = msg.Height - 7
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			d.conn.Close()
			return d, tea.Quit
		case "enter":
			cmds = append(cmds, d.sendDanmakuMsg(d.room, []byte(d.inputBar.GetValue())))
			d.inputBar.SetValue("")
		}
	case sendDanmakuRequest:
		if msg.Err != nil {
			d.insertDanmaku("NOTICE", msg.Err.Error())
		}
	case utils.ClientCreateMsg:
		d.conn = msg.Conn

		auth_msg := utils.AuthMsg{
			UID:        d.uid,
			RoomID:     d.room,
			Buvid:      d.buvid,
			Protover:   3,
			SupportAck: true,
			Scene:      "room",
			Platform:   "web",
			Key:        d.key,
			Type:       2,
		}

		cmds = append(cmds, sendAuthPkg(d.conn, auth_msg))
		return d, tea.Batch(cmds...)
	case sendAuthMsg:
		if msg.Err != nil {
			d.insertDanmaku("ERROR", msg.Err.Error())
		}

		cmds = append(cmds, receiveWsMsg(d.conn))
		cmds = append(cmds, sendHeartBeat(d.conn))

		var cmd tea.Cmd
		d.inputBar, cmd = d.inputBar.Update(msg)
		cmds = append(cmds, cmd)
		return d, tea.Batch(cmds...)
	case heartBeatMsg:
		cmds = append(cmds, sendHeartBeat(d.conn))
	case receivedDanmakuMsgs:
		for i := range msg.DanmakuMsgs {
			if msg.DanmakuMsgs[i].Err == nil {
				username, message, err := marshalMsg([]byte(msg.DanmakuMsgs[i].DanmakuMsg))
				if err == nil {
					d.insertDanmaku(string(username), string(message))
				}
			}
		}
		cmds = append(cmds, receiveWsMsg(d.conn))
		return d, tea.Batch(cmds...)
	}

	var cmd tea.Cmd
	d.inputBar, cmd = d.inputBar.Update(msg)
	cmds = append(cmds, cmd)

	return d, tea.Batch(cmds...)
}

func (d DanmakuPageModel) View() string {
	var b strings.Builder

	for range d.height - 7 - len(d.danmakus) {
		b.WriteString("\n")
	}

	for i := range d.danmakus {
		b.WriteString(d.danmakus[i].View())
	}

	b.WriteString("\n")
	b.WriteString(d.inputBar.View())
	b.WriteString("\n\n")
	b.WriteString(d.help.View(danmakuPageKeys))

	return screenStyle.Width(d.width).Height(d.height).Render(b.String())
}

func (d *DanmakuPageModel) insertDanmaku(username string, context string) {
	d.adaptiveDanmakuLength(1)
	d.danmakus = append(d.danmakus, components.NewDanmaku(username, context))
}

func (d *DanmakuPageModel) adaptiveDanmakuLength(count int) {
	if len(d.danmakus) == d.danmakuslength {
		d.danmakus = d.danmakus[count:]
	}
}

func NewDanmakuPageModel(
	uid int,
	room int,
	host string,
	buvid string,
	key string,
	bili_jct string,
	session_data string,
) DanmakuPageModel {
	d := DanmakuPageModel{
		help:           help.New(),
		danmakus:       []components.DanmakuModel{},
		danmakuslength: 0,
		inputBar:       components.NewInputModel("Comment", "Input your comment"),
		uid:            uid,
		room:           room,
		host:           host,
		buvid:          buvid,
		key:            key,
		bili_jct:       bili_jct,
		session_data:   session_data,
	}
	d.inputBar.Focus()
	return d
}
