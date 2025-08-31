package utils

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/browserutils/kooky"
	"github.com/browserutils/kooky/browser/chrome"
	"github.com/browserutils/kooky/browser/firefox"
)

func ReadCookies(path string, cookies *map[string]string) error {
	browser := ""

	switch path {
	case "firefox":
		browser = path
		path = getFirefoxCookiesFilePath()
	case "chrome", "chromium", "edge":
		browser = path
		path = getChromiumSeriesCookiesFilePath(path)
	}

	if browser == "" {
		pattern := `(firefox|chrome|chromium|edge)`
		re, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}

		browser = re.FindString(path)
		if browser == "" {
			return errors.New("This browser is not supported")
		}
	}

	var cookiesSeq kooky.CookieSeq
	switch browser {
	case "firefox":
		cookiesSeq = firefox.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("buvid3")).OnlyCookies()
		buvid := getCookies(cookiesSeq)
		cookiesSeq = firefox.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("SESSDATA")).OnlyCookies()
		session_data := getCookies(cookiesSeq)
		cookiesSeq = firefox.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("DedeUserID")).OnlyCookies()
		dede_user_id := getCookies(cookiesSeq)
		(*cookies)["buvid3"] = buvid.Value
		(*cookies)["SESSDATA"] = session_data.Value
		(*cookies)["DedeUserID"] = dede_user_id.Value

	case "chrome", "chromium", "edge":
		cookiesSeq = chrome.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("buvid3")).OnlyCookies()
		buvid := getCookies(cookiesSeq)
		cookiesSeq = chrome.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("SESSDATA")).OnlyCookies()
		session_data := getCookies(cookiesSeq)
		cookiesSeq = chrome.TraverseCookies(path, kooky.DomainHasSuffix("bilibili.com"), kooky.Name("DedeUserID")).OnlyCookies()
		dede_user_id := getCookies(cookiesSeq)
		(*cookies)["buvid3"] = buvid.Value
		(*cookies)["SESSDATA"] = session_data.Value
		(*cookies)["DedeUserID"] = dede_user_id.Value
	}

	return nil
}

func getFirefoxCookiesFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	var dir string
	switch runtime.GOOS {
	case "windows":
		dir = filepath.Join(home, "AppData", "Local", "Mozilla", "Firefox", "Profiles")
	case "darwin":
		dir = filepath.Join(home, "Library", "Application Support", "Mozilla", "Firefox", "Profiles")
	case "linux":
		dir = filepath.Join(home, ".mozilla", "firefox")
	}

	pattern := `^[a-zA-Z0-9]+\.default-release$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}
	entires, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}
	for _, entry := range entires {
		if entry.IsDir() {
			name := entry.Name()
			if re.MatchString(name) {
				return filepath.Join(dir, name, "cookies.sqlite")
			}
		}
	}
	return ""
}

func getChromiumSeriesCookiesFilePath(browser string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	var dir string
	platform := runtime.GOOS
	switch platform {
	case "windows":
		dir = filepath.Join(home, "AppData", "Local")
	case "darwin":
		dir = filepath.Join(home, "Library", "Application Support")
	case "linux":
		dir = filepath.Join(home, ".config")
	}

	switch browser {
	case "chrome":
		if platform == "linux" {
			dir = filepath.Join(dir, "google-chrome", "Default")
		} else {
			dir = filepath.Join(dir, "Google", "Chrome", "User Data", "Default", "Network")
		}
	case "chromium":
		if platform == "linux" {
			dir = filepath.Join(dir, "chromium", "Default")
		} else {
			dir = filepath.Join(dir, "Chromium", "User Data", "Default", "Network")
		}
	case "edge":
		if platform == "linux" {
			dir = filepath.Join(dir, "microsoft-edge", "Default")
		} else {
			dir = filepath.Join(dir, "Microsoft", "Edge", "User Data", "Default", "Network")
		}
	}

	dir = filepath.Join(dir, "Cookies")

	return dir
}

func getCookies(cookieSeq kooky.CookieSeq) kooky.Cookie {
	var cookies []kooky.Cookie
	for cookie := range cookieSeq {
		cookies = append(cookies, *cookie)
	}
	return cookies[0]
}

// func getAllCookies() {
// 	cookies := kooky.ReadAllCookies()
// 	for _, cookie := range cookies {
// 		fmt.Println(cookie.Name, cookie.Value)
// 	}
// }
