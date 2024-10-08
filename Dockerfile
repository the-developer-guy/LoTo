FROM golang:1.23.2-alpine3.20

COPY go.mod ./

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]