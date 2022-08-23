FROM golang:1.17-buster
ARG PROXY
ENV http_proxy ${PROXY}
ENV https_proxy ${PROXY}
ENV GO111MODULE=on
RUN mkdir backend
ADD . /backend/
WORKDIR /backend
RUN go env -w GOPROXY=https://goproxy.cn
RUN go build
EXPOSE 8080
CMD ["./news-api"]
