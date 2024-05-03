# building the go service

FROM golang:1.20 as builder
WORKDIR /application/monitoring-go-app
COPY . /application/monitoring-go-app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/monitoring-go-app -mod vendor ./main.go

# copying builds to final
FROM alpine:3.10.2 as deploy
WORKDIR /application/

COPY --from=builder /application/monitoring-go-app/build/monitoring-go-app /application/

CMD [ "sh", "-c", "/application/monitoring-go-app" ]
