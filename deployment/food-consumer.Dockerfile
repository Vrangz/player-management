FROM golang:1.19.4

RUN mkdir /app

COPY ./player-manager /app/player-manager

WORKDIR /app/player-manager/

RUN go test --race /app/player-manager/...

RUN go build -o food-consumer-job ./cmd/food-consumer/main.go

CMD [ "./food-consumer-job" ] 
