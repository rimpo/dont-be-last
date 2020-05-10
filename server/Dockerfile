FROM golang:1.14.2

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o server

CMD [ "/app/server" ]
