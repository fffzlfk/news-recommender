FROM golang:1.17.3

ENV GO111MODULE=on

RUN mkdir backend

ADD . /backend/

WORKDIR /backend

RUN go env -w GOPROXY=https://goproxy.cn

RUN go build

EXPOSE 8000

CMD ["./news-api"]