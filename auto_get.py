import os
import time

#杀死进程
def kill(pid):
	a = os.system('sudo kill -9 ' + str(pid))

#匹配进程的PID
def kill_target():
	#shell_run = 'sudo ps aux | grep {}'.format(target)
	#out = os.popen(shell_run).read()
	#return out
	shell_run = 'sudo ps -ef | grep "python3 comments.py" | grep -v "grep" | awk \'{print $2}\''
	out = os.popen(shell_run).read()
	pid = list(map(int, out.split()))
	return pid
	#pid = int(out.split()[1])
	#kill(pid)
	#pid_sudo = int(out.split()[0])
	#kill(pid_sudo)

#每10分钟重启一次脚本防止访问限制
while True :
	localtime = time.localtime(time.time())
	minute_begin = localtime.tm_min;
	os.system('sudo python3 comments.py &')
	time.sleep(600)
	p = kill_target()
	print(p)
	for PID in p:
		kill(PID)
		print("killed " + str(PID))
	print('reconnect')