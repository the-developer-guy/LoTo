FROM golang:1.23.2-alpine3.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -v -o /usr/local/bin/ ./...

CMD ["LoTo"]