sudo apt update
sudo apt install \
	cmake \
	make \
	protobuf-compiler \
	build-essential \
	autoconf \
	libtool \
	pkg-config \
	-y

PROJECT_ROOT_DIR=$(cd -- $(dirname -- $0) &> /dev/null && pwd )
pushd ${PROJECT_ROOT_DIR} || exit
export GRPC_INSTALL_DIR=${PROJECT_ROOT_DIR}/../grpc
mkdir -p ${GRPC_INSTALL_DIR}
mkdir -p ${PROJECT_ROOT_DIR}/../out/cmake/grpc-build
cmake -S ThirdParty/grpc \
    -B ../out/cmake/grpc-build \
    -DgRPC_INSTALL=ON \
    -DgRPC_BUILD_TESTS=OFF \
    -DCMAKE_INSTALL_PREFIX=${GRPC_INSTALL_DIR}
make -j 10 -C ${PROJECT_ROOT_DIR}/../out/cmake/grpc-build
make install -C ${PROJECT_ROOT_DIR}/../out/cmake/grpc-build

popd
