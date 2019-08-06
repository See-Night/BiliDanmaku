#!/bin/sh

parameter="$1"

start="start"
stop="stop"
restart="restart"

if [ "$parameter" = "$start" ];
	then
	sudo python3 auto_get.py
elif [ "$parameter" = "$stop" ];
	then
	PID=`sudo ps -ef | grep "python3 comments.py" | grep -v "grep" | awk '{print $2}'`
	sudo kill -9 $PID
	APID=`sudo ps -ef | grep "python3 auto_get.py" | grep -v "grep" | awk '{print $2}'`
	sudo kill -9 $APID
elif [ "$parameter" = "$restart" ];
	then
	PID=`sudo ps -ef | grep "python3 comments.py" | grep -v "grep" | awk '{print $2}'`
	sudo kill -9 $PID
	APID=`sudo ps -ef | grep "python3 auto_get.py" | grep -v "grep" | awk '{print $2}'`
	sudo kill -9 $APID
	sudo python3 auto_get.py
fi