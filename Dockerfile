FROM golang:1.20.3

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/

COPY . /usr/src/app/
RUN go mod tidy
RUN go build -o main cmd/api/main.go

CMD ["./main"]