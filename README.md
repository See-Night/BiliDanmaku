# BiliDanmaku

BiliDanmaku tui version. This project based bubbletea.

![screenshot](./docs/screenshot.png)

## Install

You can Download the latest release from [GitHub Releases](https://github.com/See-Night/BiliDanmaku/releases/latest).

Or build from source:

```shell
git clone -b tui https://github.com/See-Night/BiliDanmaku.git && cd BiliDanmaku
go mod tidy
chmod a+x ./build.sh
./build.sh
```

## Usage

You can get help information by executing the following command:

```shell
./BiliDanmaku -h

Bili Danmaku 0.0.1 dev Copyright © 2024-2025

Usage:
    bilidanmaku -r <room_id> -c <cookies_file>
    bilidanmaku -r <room_id> -c <browser>
    bilidanmaku -h

Parameters:
    -r <room_id>      Room ID
    -c <cookie>       Bilibili cookies
                      Please enter browser's name or browser's cookies file path
    -h                Print help
```

The operation of BiliDanmaku requires cookies, you can set cookies manually, or set `-c` parameter to enable BiliDanmaku automatically obtain cookies from browser or cookies file.

### Auto get cookie from browser

```shell
./BiliDanmaku -r <room_id> -c <browser>
```

`<browser>` can replaced by `firefox`, `chrome`, `edge` and `chromium` now.

### Set cookie manually

```shell
./Bilidanmaku -r <room_id> -u <uid> -s <session_data> -b <buvid3>
```

You can find `uid`, `session_data` and `buvid3` from browsers.
