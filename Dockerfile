FROM golang:1.15.2-buster

WORKDIR /app

COPY . .

ENV PORT=9090

EXPOSE 9090

RUN ["go", "install"]
RUN ["go", "build"]
RUN ["mv","zi","/usr/bin"]
RUN ["wget", "https://github.com/fiorix/go-daemon/releases/download/v1.3/go-daemon_1.3_amd64.deb"]
RUN ["apt","install", "./go-daemon_1.3_amd64.deb"]
RUN ["rm", "./go-daemon_1.3_amd64.deb"]

CMD [ "zi", "serve", "9090" ]