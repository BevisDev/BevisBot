FROM golang:1.25.5 AS builder

RUN apt-get update &&  \
    apt-get install -y --no-install-recommends \
    make \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM ubuntu:22.04

ENV TZ=Asia/Ho_Chi_Minh

RUN apt-get update &&  \
    apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /build/config ./config
COPY --from=builder /build/bevis ./

ENTRYPOINT [ "./bevis" ]