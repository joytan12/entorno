FROM golang:1.23.1 AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o financiero .

WORKDIR /app

CMD ["./financiero"]