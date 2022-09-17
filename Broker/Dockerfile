FROM golang:1.19-alpine

WORKDIR /broker

COPY . .

RUN go mod download

RUN go build -o ./broker ./cmd/app

EXPOSE 43200 43201

CMD ["/broker/broker"]