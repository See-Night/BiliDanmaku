# CommentsLogByMySQL
用MySQL来储存记录直播间弹幕信息，供房管查询使用
### 通过python脚本将直播间弹幕扒下来存在数据库中，**需要MySQL支持**，已实装B站直播自动检测功能

***仅支持BILIBILI直播***

身为一名房管（DD），平时总有很多人私信询问“啊，我只在直播间发了一个“xxx”怎么就被封了，是不是误判了”  
一方面为了给大家一个公道  
另一方面也是防止有人钻空子  
再有就是减轻房管的压力  
我，DD（X），用业余时间写了这么一个小东西，方便弹幕记录  
  
***由于本身没学过python所以相当于现学现用，代码写得辣鸡的一批***  

**由于我本身是mea gachi所以所有的配置是以mea直播间为范本的，如需在其他直播间使用请自行更改脚本，需要更改的位置已在脚本中指出**
  
#### 安装方法
```
apt：
$ wget https://dreammer12138.github.io/Documents/CommentsLogByMySQL/apt-setup.sh
$ sudo chmod 777 ./apt-setup.sh
$ ./apt-setup.sh

yum:
$ wget https://dreammer12138.github.io/Documents/CommentsLogByMySQL/yum-setup.sh
$ sudo chmod 777 ./yum-setup.sh
$ ./yum-setup.sh

```

#### 配置
```
** MySQL **  

配置MySQL用户密码
$ sudo mysql
>> use mysql;
>> update user set plugin="mysql_native_password" where user="root";
>> flush privileges;
>> update user set authentication_string=PASSWORD("新密码") where user="root";
>> flush privileges;
>> exit;

重启MySQL
$ sudo /etc/init.d/mysql restart

登录MySQL（不使用sudo权限）,如果可以登录进去就说明配置完成
$ mysql -u root -p

新建数据库
>> create database Bili_Comments character set 'utf8' collate 'utf8_general_ci';
>> exit;

** Python **

$ sudo vim /usr/CommentsLogByMySQL/comments.py

将数据库信息中的密码改成你刚才设置的密码，保存退出

	conn = pymysql.connect(
			host = 'localhost',
			port = 3306,
			user = 'root',           #MySQL用户名
			password = '123456',     #MySQL密码 << 改这个，注意不要忘了用引号
			db = 'Bili_Comments',    #数据库名称
			charset = 'utf8'
		)

```

安装位置：`/usr/CommentsLogByMySQL/`  

#### 使用方法
```
$ sudo bash /usr/CommentsLogByMySQL/CommentsLogByMySQL.sh
```
建议使用screen单独运行  

创建新进程 `$ screen -S <你的进程名称>`  
后台运行进程 `Ctrl + A + D`  
返回进程 `$ screen -R <你的进程名称>`  

#### 自定义直播间部署
自定义直播间需要更改`comments.py`和`check.py`两个文件  

*comments.py*

```
form_data = {
		"roomid": "12235923",         << 改这里，roomid即为直播间号码
		"csrf_token": "",
		"csrf": "",
		"visit_id": ""
	}
```

*check.py*

```
headers['Referer'] = 'https://space.bilibili.com/349991143?from=search&seid=16603871590950900377'

req = requests.get("https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld?mid=349991143")
```

打开个人空间  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151510.png)  

F12打开资源管理器，选择上面Network选项卡，F5刷新，搜索getRoom  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151616.png)  

单击搜索结果，在详细信息中找到Request Headers -> Referer，然后将上面那行代码内容替换  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190810101958.png)

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151635.png)  
