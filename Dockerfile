FROM golang:1.15.2-buster

WORKDIR /app

COPY . .

ENV PORT=9090

EXPOSE 9090

WORKDIR /app/man
RUN ["go", "install"]
RUN ["go", "build", "-o", "../bin/zi"]
WORKDIR /app/bin
RUN ["mv","zi","/usr/bin"]
WORKDIR /app
RUN ["wget", "https://github.com/fiorix/go-daemon/releases/download/v1.3/go-daemon_1.3_amd64.deb"]
RUN ["apt","install", "./go-daemon_1.3_amd64.deb"]
RUN ["rm", "./go-daemon_1.3_amd64.deb"]

CMD [ "zi", "serve", "9090" ]