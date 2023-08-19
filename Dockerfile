FROM golang:1.20 AS buildenv

LABEL maintainer="better.tian@qq.com"

ARG SERVICE
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# Create a location in the container for the source code.
RUN mkdir -p /app

WORKDIR /app
COPY go.* /app/
COPY script/start.sh /app
# COPY vendor vendor
COPY pkg pkg
COPY app/grpc-${SERVICE} app/grpc-${SERVICE}

RUN go mod download
RUN go mod verify

# RUN make gen service=${SERVICE} 
COPY kitex_gen kitex_gen
RUN go build \
    -mod=readonly \
    -ldflags="-w -s" \
    -a -o ${SERVICE} app/grpc-${SERVICE}/main.go

FROM alpine:3.16.2 AS runnerenv
RUN mkdir -p /app
WORKDIR /app
ARG SERVICE
ARG VERSION
ARG GIT_COMMIT

ENV SERVICE=${SERVICE}
ENV VERSION=${VERSION}
ENV GIT_COMMIT=${GIT_COMMIT}

RUN echo "https://mirrors.aliyun.com/alpine/v3.16/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.16/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

COPY --from=buildenv /app/${SERVICE}  /app/start.sh /app
# COPY --from=buildenv /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["/app/start.sh"]
