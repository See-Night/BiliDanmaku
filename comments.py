import requests
import pymysql
import datetime
import re

cache = {}

#检查表是否存在，存在则返回1，不存在则返回0
def table_exists(con, table_name):
	sql = "SHOW TABLES;"
	con.execute(sql)
	tables = [con.fetchall()]
	table_list = re.findall('(\'.*?\')', str(tables))
	table_list = [re.sub("'", '', each) for each in table_list]
	if table_name in table_list:
		return 1
	else:
		return 0

while  True:
	#数据库连接信息
	conn = pymysql.connect(
		host = '127.0.0.1',
		port = 3306,
		user = 'root',
		password = '123456',
		db = 'Bili_Comments',
		charset = 'utf8'
	)

	cursor = conn.cursor()

	#直播间信息，csrf和csrf_token每7天会更新，需要手动更换
	form_data = {
		"roomid": "12235923",
		"csrf_token": "",
		"csrf": "",
		"visit_id": ""
	}

	#post请求中的headers
	headers = {}
	headers['Contect-Type'] = 'application/x-www-form-urlencoded'
	headers['User-Agent'] = 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36'
	headers['Connection'] = 'close'

	try:
		#post请求
		res = requests.post("https://api.live.bilibili.com/ajax/msg", headers = headers, data = form_data)
		#获取直播间弹幕
		result = res.json()['data']['room']
		#获取当前时间，用来创建表
		now_datetime = datetime.datetime.now().strftime('%Y%m%d')
		if(table_exists(cursor, 'Comments' + now_datetime) != 1):
			SQL = 'Create table Comments' + now_datetime + ' ( Nickname VARCHAR(255) NOT NULL, UID INT NOT NULL, date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, text VARCHAR(255) NOT NULL, ct VARCHAR(255) NOT NULL, ts VARCHAR(255) NOT NULL);'
			cursor.execute(SQL)
		#匹配弹幕信息，将新的弹幕插入到表中
		for comments_msg in result:
			SQL = 'select ts from Comments' + str(comments_msg['timeline'][0:10].replace('-', '')) + ' where ts="' + str(comments_msg['check_info']['ts']) + '";'
			cursor.execute(SQL)
			res_select = cursor.fetchall()
			if res_select != ():
				pass
			else:
				SQL = 'insert into Comments' + comments_msg['timeline'][0:10].replace('-', '') + '(Nickname, UID, text, ct, ts) VALUES("' + comments_msg['nickname'] + '",' + '%d'%comments_msg['uid'] + ',"' + comments_msg['text'] + '","' + comments_msg['check_info']['ct'] + '","' + '%d'%comments_msg['check_info']['ts'] + '");'
				cursor.execute(SQL)
				conn.commit()
				print(comments_msg['nickname'] + ':' + comments_msg['text'])
		#缓存池匹配弹幕，由于上面使用了直接从数据库中查询匹配的方法，所以将此方法注释
		#cursor.execute("select * from Comments" + now_datetime + " order by ts desc limit 1;")
		#res_select = cursor.fetchone()
		#ct = str(res_select[4])
		#ts = int(res_select[5])
		#cache['ts'] = int(ts)
		#cache['ct'] = str(ct)

		#if (cache['ct'] == result['check_info']['ct']) and (cache['ts'] == result['check_info']['ts']):
		#	pass
		#else:
		#	cache = result['check_info']
		#	SQL = 'insert into Comments' + result['timeline'][0:10].replace('-', '') + '(Nickname, UID, text, ct, ts) VALUES("' + result['nickname'] + '",' + '%d'%result['uid'] + ',"' + result['text'] + '","' + result['check_info']['ct'] + '","' + '%d'%result['check_info']['ts'] + '");'
		#	cursor.execute(SQL)
		#	conn.commit()
		#	print(result['nickname'] + ':' + result['text'])
		s = requests.session()
		s.keep_alive = False
		cursor.close()
		conn.close()
	except:
		print("Error has been ignored")
		continue
