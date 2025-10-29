FROM golang:1.25.3 AS builder AS builder

RUN apt-get update &&  \
    apt-get install -y --no-install-recommends \
    libaio1 \
    make \
    unzip && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /build

COPY ./src/resources/oracle /oracle/

RUN unzip /oracle/instantclient-basic-linux*.zip -d /oracle && \
    rm /oracle/instantclient-basic-linux*.zip

ENV LD_LIBRARY_PATH=/oracle/instantclient_12_2 \
    ORACLE_HOME=/oracle/instantclient_12_2 \
    CGO_ENABLED=1

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM ubuntu:22.04

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Ho_Chi_Minh

RUN apt-get update &&  \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    libaio1 \
    tzdata && \
    echo "$TZ" > /etc/timezone && \
    ln -snf "/usr/share/zoneinfo/${TZ}" /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /oracle /oracle

WORKDIR /app

ENV LD_LIBRARY_PATH=/oracle/instantclient_12_2 \
    ORACLE_HOME=/oracle/instantclient_12_2 \
    CGO_ENABLED=1

COPY --from=builder /build/config/prod.yml ./config/prod.yml
COPY --from=builder /build/bevis ./

ENTRYPOINT [ "./bevis" ]