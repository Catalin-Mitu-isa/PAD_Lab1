FROM golang

WORKDIR /broker

COPY . .

RUN apt-get update && apt-get install make protobuf-compiler protobuf-compiler-grpc golang -y

ENV CGO_ENABLE=0
ENV GOPATH=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin

RUN make

EXPOSE 43200 43201

CMD ["make", "run"]