FROM golang:1.18-alpine

COPY go.mod ./
COPY go.sum ./
ENV GOPATH=/

RUN go mod download

COPY ./ ./

RUN go build -o transactions-app ./cmd/main.go

EXPOSE 8080

CMD [ "./transactions-app"]