FROM golang:1.15

COPY ./server.go .

RUN go build -o ./server.go

CMD ./server.go
EXPOSE 8000
