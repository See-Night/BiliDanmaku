#!/bin/sh

liveon="正在直播"
liveoff="直播停止"
live=0

accpet=`cat ./config.json | jq -r '.getInfo.headers.Accept'`
origin=`cat ./config.json | jq -r '.getInfo.headers.Origin'`
referer=`cat ./config.json | jq -r '.getInfo.headers.Referer'`
useragent=`cat ./config.json | jq -r '.getInfo.headers."User-Agent"'`
url=`cat ./config.json | jq -r '.getInfo.url'`
roomid=`cat ./config.json | jq -r '.comments.room.roomid'`

var=`curl -s -H "$accpet" -H "$origin" -H "$referer" -H "useragent" $url | jq -r '.data.name'`
echo "正在监控$var的直播间"
echo "直播间号：$roomid"

while :
do
	MSG=`sudo python3 /usr/CommentsLogByMySQL/check.py`
	if [ $MSG = $liveon ]
		then
		if [ $live == 0 ]
			then
			sudo ./auto_get.sh start &
			live=1
		fi
	elif [ $MSG = $liveoff ]
		then
		if [ $live == 1 ]
			then
			sudo ./auto_get.sh stop &
			live=0
		fi
	fi
done