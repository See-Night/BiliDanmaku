#!/usr/bin/python3
from bilibili_api.live import LiveDanmaku, get_room_info, get_room_play_info
import sys
from openpyxl import Workbook
import json
import time

class DMK:
    def __init__(self, roomid):
        self.roomid = roomid
        self.realid = get_room_play_info(roomid)['room_id']
        self.title = get_room_info(self.realid)['room_info']['title']
        self.excel = Workbook()
        sheet = self.excel.active
        sheet['A1'] = 'user'
        sheet['B1'] = 'uid'
        sheet['C1'] = 'timestamp'
        sheet['D1'] = 'content'

    def push(self, danmaku):
        sheet = self.excel.active
        sheet.append(danmaku)

    def close(self):
        self.excel.save(
            '{title}-{roomid}.xlsx'.format(
                title=self.title, 
                roomid=self.roomid
            )
        )

def main():
    try:
        roomid = int(sys.argv[1])
    except Exception:
        print("请输入直播间号码")
        return

    room = LiveDanmaku(room_display_id=roomid)
    dmk = DMK(roomid)
    
    @room.on("DANMU_MSG")
    async def on_danmu(msg):
        user = msg['data']['info'][2][1]
        uid = msg['data']['info'][2][0]
        timestamp = time.strftime(
            '%Y-%m-%d %H:%M:%S',
            time.localtime(msg['data']['info'][0][4]/1000)
        )
        content = msg['data']['info'][1]
        print(
            '[{roomid}][{user}][{uid}][{timestamp}]{content}'
            .format(
                roomid=roomid,
                user=user,
                uid=uid,
                timestamp=timestamp,
                content=content
            )
        )
        dmk.push([user, uid, timestamp, content])
    
    try:
        room.connect()
    except KeyboardInterrupt:
        print('结束')
    finally:
        dmk.close()

if __name__ == "__main__":
    main()
