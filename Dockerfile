FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add python3 py3-pip gcc g++ python3-dev tzdata

RUN ls /usr/share/zoneinfo

RUN date

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /app

COPY . .

RUN pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

ENTRYPOINT ["python3", "app.py"]
