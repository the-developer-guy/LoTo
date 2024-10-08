FROM golang:1.23.2-alpine3.20

WORKDIR /usr/src/app

COPY go.mod ./

COPY ./ ./
RUN go build -v -o /usr/local/bin/ ./...

CMD ["LoTo"]