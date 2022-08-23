FROM golang:1.17-buster
ARG PROXY
ENV PROXY ${PROXY}
ENV GO111MODULE=on
RUN mkdir backend
ADD . /backend/
WORKDIR /backend
RUN go env -w GOPROXY=https://goproxy.cn
RUN go build
EXPOSE 8080
CMD ["./news-api"]
