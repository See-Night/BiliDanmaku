#!/usr/bin/python3

import numpy as np
from ws4py.client.threadedclient import WebSocketClient
import zlib
import json
import re
import time
import sys

def writeInt(buf, start, length, value):
    for i in range(0, length):
        buf[start + i] = value/(pow(256, length - i - 1))
    return buf

def encode(jsonstr, op):
    data = np.frombuffer(jsonstr, dtype=np.uint8)
    packetLen = 16 + len(data)
    header = [0, 0, 0, 0, 0, 16, 0, 1, 0, 0, 0, op, 0, 0, 0, 1]
    header = writeInt(header, 0, 4, packetLen)
    res = np.array(np.append(header, data), dtype='uint8')
    return res.tobytes()

def readInt(array):
    return int.from_bytes(array, byteorder='big')

def decode(blob):
    packetLen = readInt(blob[0:4])
    headerLen = readInt(blob[4:6])
    ver = readInt(blob[6:8])
    op = readInt(blob[8:12])
    seq = readInt(blob[12:16])
    body = []
    if op == 5:
        try:
            temp = zlib.decompress(blob[16:])[16:].decode('utf-8')
            group = re.split('[\x00-\x1f]+', temp)
            for i in group:
                try:
                    body.append(json.loads(i))
                except Exception:
                    pass
        except zlib.error:
            pass
        except UnicodeDecodeError:
            pass
    elif op == 3:
        body = {
            "count": blob[16:20]
        }
    return {
        'packetLen': packetLen,
        'headerLen': headerLen,
        'ver': ver,
        'op': op,
        'seq': seq,
        'body': body
    }

class bilichat(WebSocketClient):
    def opened(self):
        self.send(encode(bytes('{"roomid":' + self.roomid + '}', encoding="utf8"), 7))
    def closed(self, code, reason=None):
        print(code)
    def received_message(self, m):
        packet = decode(m.data)
        op = packet['op']
        if op == 8:
            print('加入房间')
        elif op == 5:
            for i in packet['body']:
                try:
                    if i['cmd'] == 'DANMU_MSG':
                        print('{}: {}'.format(i['info'][2][1], i['info'][1]))
                    else:
                        pass
                except TypeError:
                    continue
        else:
            pass

try:
    if len(sys.argv) < 2:
        raise Exception("请输入直播间地址")
    ws = bilichat('wss://broadcastlv.chat.bilibili.com/sub')
    ws.roomid = sys.argv[1]
    print(ws.roomid)
    ws.connect()
    while True:
        time.sleep(30)
        ws.send(encode(b'', 2))
    ws.run_forever()
except KeyboardInterrupt:
    print('退出')
except Exception as e:
    print(e)
