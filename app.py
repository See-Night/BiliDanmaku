#!/usr/bin/python3
import sys
import subprocess
import time
import os
import signal
from bilibili_api import live

def app():
    try:
        roomid = sys.argv[1]
        ison = False

        while True:
            live_status = live.get_room_play_info(roomid)['live_status']
            if live_status == 1 and ison == False:
                ison = True
                print('[{roomid}][{time}][INFO] 直播开始'.format(
                    roomid=roomid,
                    time=time.strftime(
                        '%Y-%m-%d %H-%M-%S',
                        time.localtime()
                    )
                ))
                p = subprocess.Popen(['python3', './BiliDanmaku.py', '{roomid}'.format(roomid=roomid)], shell=False)
            elif live_status in [0, 2] and ison == True:
                p.send_signal(signal.SIGINT)
                ison = False
            time.sleep(120)
    except Exception:
        return
    except KeyboardInterrupt:
        return

if __name__ == "__main__":
    app()
