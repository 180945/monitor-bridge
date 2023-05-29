FROM golang:1.19

RUN apt-get update && apt-get upgrade -y

RUN apt install build-essential -y

RUN git clone --depth 1 https://github.com/180945/monitor-bridge

WORKDIR /go/monitor-bridge

RUN go build -o bridge_monitor

CMD ["./bridge_monitor"]