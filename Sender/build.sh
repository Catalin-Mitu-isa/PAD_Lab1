#!/bin/bash

sudo apt update
sudo apt install \
	cmake \
	make \
	g++ \
	libgtk-3-dev
	build-essential \
	pkg-config \

	autoconf \
	libtool \
	-y

mkdir -p ./cmake-build

cmake -S . -B ./cmake-build
make ./cmake-build
