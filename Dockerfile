FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN apk add build-base

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server .

EXPOSE 8080

CMD ["./server"]
