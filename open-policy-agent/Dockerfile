# building the go service

FROM golang:1.19 as builder
WORKDIR /application/open-policy-agent
COPY . /application/open-policy-agent

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/open-policy-agent -mod vendor ./main.go

# copying builds to final
FROM alpine:3.10.2 as deploy
WORKDIR /application/

COPY --from=builder /application/open-policy-agent/build/open-policy-agent /application/

CMD [ "sh", "-c", "/application/open-policy-agent" ]
