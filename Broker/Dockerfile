FROM golang:1.19-alpine

WORKDIR /broker

COPY . .

RUN make

EXPOSE 43200 43201

CMD ["make", "run"]