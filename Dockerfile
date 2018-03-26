FROM golang:1.10 as go
WORKDIR /go/src/github.com/DanShu93/sleeper
COPY . .

RUN CGO_ENABLED=0 go build server.go

FROM alpine

WORKDIR /app

COPY --from=go /go/src/github.com/DanShu93/sleeper/server .

ARG MAX

CMD /app/server --max $MAX
