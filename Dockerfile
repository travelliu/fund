# Builder
FROM golang:1.13 as builder
# change china 163 source
RUN sed -i "s/deb.debian.org/mirrors.163.com/g" /etc/apt/sources.list && sed -i "s/security.debian.org/mirrors.163.com/g" /etc/apt/sources.list
RUN apt-get -o Acquire::Check-Valid-Until=false update \
    && apt-get upgrade -y \
    && apt-get install -y git gcc make
WORKDIR /go/src/github.com/travelliu/fund
COPY . .
#ENV GO111MODULE on
RUN make build


# Distribution
FROM debian:buster
RUN sed -i "s/deb.debian.org/mirrors.163.com/g" /etc/apt/sources.list && sed -i "s/security.debian.org/mirrors.163.com/g" /etc/apt/sources.list
RUN apt-get -o Acquire::Check-Valid-Until=false update \
    && apt-get upgrade -y \
    && apt-get install -y  tzdata ca-certificates \
    # for SAML
    xmlsec1 \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo Asia/Shanghai > /etc/timezone \
    && update-ca-certificates 2>/dev/null || true
WORKDIR /app
EXPOSE 8081
COPY --from=builder /go/src/github.com/travelliu/fund/fund /app
COPY conf.yaml  /app
ENTRYPOINT /app/fund