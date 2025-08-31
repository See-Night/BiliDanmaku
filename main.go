package main

import (
	"bilidanmaku/ui/pages"
	"bilidanmaku/utils"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

var VERSION = "0.0.1 dev"

func printHelp() {
	fmt.Printf("Bili Danmaku %s Copyright © 2024-2025\n", VERSION)
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    bilidanmaku -r <room_id> -c <cookies_file>")
	fmt.Println("    bilidanmaku -r <room_id> -c <browser>")
	fmt.Println("    bilidanmaku -h")
	fmt.Println("")
	fmt.Println("Parameters:")
	fmt.Println("    -r <room_id>      Room ID")
	fmt.Println("    -c <cookie>       Bilibili cookies")
	fmt.Println("                      Please enter browser's name or browser's cookies file path")
	fmt.Println("    -h                Print help")
}

func main() {
	var room_id = flag.Int("r", 0, "Room ID")
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

	var uid int
	var buvid string
	var bili_jct string

	cookies := make(map[string]string)
	if *cookie != "" {
		// Get cookies from browser
		utils.ReadCookies(*cookie, &cookies)

		uid, _ = strconv.Atoi(cookies["DedeUserID"])
		buvid = cookies["buvid3"]
		bili_jct = cookies["bili_jct"]
	}

	true_id, host, key, err := utils.GetWebsocketConf(strconv.Itoa(*room_id), &(map[string]string{
		"SESSDATA": cookies["SESSDATA"],
		"buvid3":   cookies["buvid3"],
	}))
	if err != nil {
		return
	}

	// Enable debug log
	// f, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	fmt.Println("fatal:", err)
	// 	os.Exit(1)
	// }
	// defer f.Close()

	m := pages.NewDanmakuPageModel(uid, true_id, host, buvid, key, bili_jct, cookies["SESSDATA"])
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
