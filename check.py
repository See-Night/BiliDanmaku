import requests
import os
#from bs4 import BeautifulSoup as bs

live = 0

def changelive(l):
	global live
	live = l

headers = {}

headers['Accept'] = 'application/json, text/plain, */*'
headers['Origin'] = 'https://space.bilibili.com'
headers['Referer'] = 'https://space.bilibili.com/349991143?from=search&seid=16603871590950900377'
headers['User-Agent'] = 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'

req = requests.get("https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld?mid=349991143")
result = req.json()['data']['liveStatus']
if result == 1:
	print("正在直播")
	#os.system("sudo ./auto_get.sh start")
	changelive(1)	
else:
	print("直播停止")
	if live == 1:
		#os.system("sudo ./auto_get.sh stop")
		changelive(0)
	else:
		pass
		#print("cccc")
