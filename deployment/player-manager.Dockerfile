FROM golang:1.19.4

RUN mkdir /app
RUN mkdir /resources

COPY ./player-manager /app/player-manager
COPY ./swagger-ui-dist /resources/swagger-ui-dist

WORKDIR /app/player-manager/

RUN go test --race /app/player-manager/...

RUN go build -o player-manager-app ./cmd/main/main.go

CMD [ "./player-manager-app" ] 
