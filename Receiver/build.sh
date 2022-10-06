#!/bin/bash

sudo apt update
sudo apt install \
	cmake \
	make \
	protobuf-compiler \
	protobuf-compiler-grpc \
	build-essential \
	autoconf \
	libtool \
	pkg-config \
	-y

mkdir -p ../../cmake/Sender/

cmake -S . -B ../../cmake/Sender
make ../../cmake/Sender