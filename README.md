# BiliDanmaku

BiliDanmaku是一个命令行工具，用于记录bilibili**直播**弹幕。

BiliDanmaku is a CLI tools for recording bilibili live Danmaku.

## Before use

在使用之前，请确保你的电脑安装了Python

Before using it, you need to make sure your device has Python installed.

### Download Code

从Github上下载源代码。

Download source code from github.

或者从github上克隆源码。

Or clone source code from github by git.

```shell
$ git clone https://github.com/See-Night/BiliDanmaku.git
```

### Install package

```shell
$ pip install -r requirements.txt
```

### Push image from docker hub

你也可以从[docker hub](https://hub.docker.com/r/seenight/bilidanmaku)上获取镜像。

You also can push image from [docker hub](https://hub.docker.com/r/seenight/bilidanmaku). 

```bash
$ docker pull seenight/bilidanmaku
```

## Usage

### Run script

```shell
$ ./BiliDanmaku.py <roomid>
```

When it is stoped, it will generated a excel file.

### Run script that automate

```bash
$ ./app.py <roomid>
```

It will record danmaku automatically.

### Run docker script

```bash
$ ./docker-start <roomid> <path to save>
```

## Thanks

Thanks for [@lovelyyoshino](https://github.com/lovelyyoshino/)'s [Bilibili-Live-API](https://github.com/lovelyyoshino/Bilibili-Live-API) and [@Passkou](https://github.com/Passkou)'s [bilibili-api](https://github.com/Passkou/bilibili-api)

