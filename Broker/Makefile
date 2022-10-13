NAME = broker
APP = ./cmd/app

${NAME}: all
	$(info [+] Go-Broker)

all: deps proto
	$(info [+] Compiling go)
	@go build -o ./${NAME} ${APP}

deps:
	$(info [+] Downloading all dependencies)
	@go mod download

proto: deps
	$(info [+] Generating proto files)
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
	$(info [+] Moving Receiver proto files to right directory)
	@mv ./proto/R*.pb* ./pkg/receiver
	$(info [+] Moving Sender proto files to right directory)
	@mv ./proto/S*.pb* ./pkg/sender
	$(info [+] Moving Broker proto files to right directory)
	@mv ./proto/B*.pb* ./pkg/broker

run: all
	$(info [+] Running brocker)
	@./${NAME}

.PHONY = deps all proto
