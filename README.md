# CommentsLogByMySQL
用MySQL来储存记录直播间弹幕信息，供房管查询使用
### 通过python脚本将直播间弹幕扒下来存在数据库中，**需要MySQL支持**

身为一名房管（DD），平时总有很多人私信询问“啊，我只在直播间发了一个“xxx”怎么就被封了，是不是误判了”  
一方面为了给大家一个公道  
另一方面也是防止有人钻空子  
再有就是减轻房管的压力  
我，DD（X），用业余时间写了这么一个小东西，方便弹幕记录  
  
***由于本身没学过python所以相当于现学现用，代码写得辣鸡的一批  
  
#### 安装方法
```
Ubuntu/Deepin/Debian：
$ wget https://dreammer12138.github.io/Documents/setup.sh
$ sudo chmod 777 ./setup.sh
$ ./setup.sh

CentOS:
$ sudo yum -y install mysql-server mysql-client
然后自行安装python3，网上教程一抓一大把这里就不详细说明了
然后安装python库
$ sudo pip3 install pymysql
$ sudo pip3 install requests
```

#### 使用方法
```
$ sudo auto_get.sh + 参数  
可选参数：  
  start     启动
  stop      停止
  restart   重启
```
建议使用screen单独运行
