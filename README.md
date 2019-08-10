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
  
安装位置：`/usr/CommentsLogByMySQL/`  

#### 配置  

** MySQL **  
```
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

```
  
** config.json **  
  
***所有的配置信息都存储在config.json文件中，修改前请务必备份好原配置***

数据库储存：修改“comments”->“mysql”  
```
"mysql": {
	"host": "localhost",		<-数据库地址，如果是本地数据库就默认不要修改
	"port": 3306,				<-数据库端口，默认3306，不建议修改
	"user": "root",				<-mysql用户名，不建议修改
	"password": "123456",		<-mysql密码，根据配置数据库时设置的密码修改
	"db": "Bili_Comments",		<-数据库名，根据配置数据库时新建的数据库修改
	"charset": "utf8"			<-数据库数据编码方式，默认utf-8，不建议修改
}
```
  
弹幕姬配置：修改“comments”->“room”
```
"room": {
	"roomid": "12235923",		<-这里填写直播间号码
	"csrf_token": "",
	"csrf": "",
	"visit_id": ""
}
```
  
直播检测配置：修改“check”和“getInfo”
```
"check": {
	"headers": {
		"Accept": "application/json, text/plain, */*",
		"Origin": "https://space.bilibili.com",
		"Referer": "https://space.bilibili.com/349991143",								<-修改此处
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	},
	"url": "https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld?mid=349991143"	<-修改此处只需要把mid后面的那串数字修改为主播个人空间的号码即可
},
"getInfo": {
	"headers": {
		"Accept": "application/json, text/plain, */*",
		"Origin": "https://space.bilibili.com",
		"Referer": "https://space.bilibili.com/349991143",								<-修改此处
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	},
	"url": "https://api.bilibili.com/x/space/acc/info?mid=349991143&jsonp=jsonp"		<-修改此处只需要把mid后面的那串数字修改为主播个人空间的号码即可
}
```
打开个人空间  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151510.png)  

F12打开资源管理器，选择上面Network选项卡，F5刷新，搜索getRoom  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151616.png)  

单击搜索结果，在详细信息中找到Request Headers -> Referer，然后将上面那行代码内容替换  

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190810101958.png)

![kaguramea](https://dreammer12138.github.io/Documents/CommentsLogByMySQL/dict/20190808151635.png)  

#### 使用方法
```
$ sudo bash /usr/CommentsLogByMySQL/CommentsLogByMySQL.sh
```
建议使用screen单独运行  

创建新进程 `$ screen -S <你的进程名称>`  
后台运行进程 `Ctrl + A + D`  
返回进程 `$ screen -R <你的进程名称>`  


