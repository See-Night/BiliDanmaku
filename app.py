#!/usr/bin/python3
import sys
import subprocess
import time
import os
import signal
import asyncio
from bilibili_api.live import LiveRoom

async def app():
    try:
        roomid = sys.argv[1]
        ison = False

        while True:
            live_status = (await LiveRoom(roomid).get_room_play_info())['live_status']
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
    except Exception as e:
        print(e)
        return
    except KeyboardInterrupt:
        return

asyncio.get_event_loop().run_until_complete(app())
