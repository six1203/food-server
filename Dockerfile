FROM golang:1.19 AS builder
COPY . /src
WORKDIR /src
ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
RUN make build


FROM ubuntu:18.04
WORKDIR /app
ENV TZ=Asia/Shanghai
COPY --from=builder /src/bin /app
RUN sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
    && sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt update \
    && apt install -y tzdata \
    && ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["./food-server", "-conf", "/data/conf"]
