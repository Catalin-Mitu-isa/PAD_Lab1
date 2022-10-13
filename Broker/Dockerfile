FROM golang

WORKDIR /broker

COPY . .

RUN apt-get update && apt-get install make protobuf-compiler protobuf-compiler-grpc golang -y

RUN make

EXPOSE 43200 43201

CMD ["make", "run"]