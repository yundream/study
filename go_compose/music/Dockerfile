FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go mod download
RUN go build -o /docker-go-music

EXPOSE 8088

CMD ["/docker-go-music"]