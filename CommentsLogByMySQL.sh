#!/bin/sh

liveon="正在直播"
liveoff="直播停止"
live=0

while :
do
	MSG=`sudo python3 ./check.py`
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