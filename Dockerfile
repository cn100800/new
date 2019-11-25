# docker run -d --name news --restart always cn100800/news
FROM alpine
MAINTAINER freecracy1024@gmail.com
ARG PASSWORD
ENV TZ Asia/Shanghai
RUN apk update && \
    apk add --no-cache tzdata libgcc libstdc++ libc6-compat  libc-dev && \
    cp -f  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk add go git && \
    GO111MODULE=on go get github.com/freecracy/news && \
    echo "58 23 * * * /root/go/bin/news -o -u myqq2018@gmail.com -P $PASSWORD -h smtp.gmail.com -f myqq2018@gmail.com -t myqq2018@gmail.com -p 587" > /root/hello-cron && \
    chmod 0644 /root/hello-cron && \
    crontab /root/hello-cron
CMD crond -f
