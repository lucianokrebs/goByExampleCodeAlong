FROM golang:1.16

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/runner

ENTRYPOINT ["./main"]