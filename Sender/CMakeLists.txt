cmake_minimum_required(VERSION 3.22)
project(Sender)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Use the package PkgConfig to detect GTK+ headers/library files
find_package(PkgConfig REQUIRED)
pkg_check_modules(GTK3 REQUIRED gtk+-x11-3.0)

# For automatic third party library setup
include(FetchContent)
set(FETCHCONTENT_QUIET OFF)

# Asio -------------------------------------------------------------------------

set(ASIO_TAG "asio-1-24-0")
set(ASIO_SRC "${PROJECT_SOURCE_DIR}/../${ASIO_TAG}-src")

FetchContent_Declare(asio
        GIT_REPOSITORY  https://github.com/chriskohlhoff/asio.git
        GIT_TAG         ${ASIO_TAG}
        GIT_SHALLOW     TRUE
        SOURCE_DIR      ${ASIO_SRC}
        GIT_PROGRESS    TRUE
        )

FetchContent_GetProperties(asio)

if(NOT asio_POPULATED)
    message(STATUS "Fetching asio ${ASIO_TAG}")

    if(EXISTS "${ASIO_SRC}/asio")
        set(FETCHCONTENT_SOURCE_DIR_ASIO ${ASIO_SRC})
        message(STATUS "asio ${ASIO_TAG} source found at ${ASIO_SRC}; skip fetching.")
    endif()

    FetchContent_Populate(asio)
    message(STATUS "Populating asio ${ASIO_TAG} done; src at ${asio_SOURCE_DIR}")
endif()

add_library(asio INTERFACE)

target_include_directories(asio
        INTERFACE ${asio_SOURCE_DIR}/asio/include
        )

target_compile_definitions(asio
        INTERFACE
        ASIO_STANDALONE
        ASIO_NO_DEPRECATED
        )

target_link_libraries(asio
        INTERFACE pthread
        )

# Asio end ---------------------------------------------------------------------

# sciter -----------------------------------------------------------------------

set(SCITER_TAG "5.0.0.9")
set(SCITER_SDK "${PROJECT_SOURCE_DIR}/../sciter-js-sdk")

FetchContent_Declare(sciter
        GIT_REPOSITORY  https://gitlab.com/sciter-engine/sciter-js-sdk.git
        GIT_TAG         ${SCITER_TAG}
        GIT_SHALLOW     TRUE
        SOURCE_DIR      ${SCITER_SDK}
        GIT_PROGRESS    TRUE
        )

FetchContent_GetProperties(sciter)

if(NOT sciter_POPULATED)
    message(STATUS "Fetching sciter ${SCITER_TAG}")

    if(EXISTS "${SCITER_SDK}")
        set(FETCHCONTENT_SOURCE_DIR_SCITER ${SCITER_SDK})
        message(STATUS "sciter ${SCITER_TAG} found; skip fetching.")
    endif()

    FetchContent_Populate(sciter)
    message(STATUS "Populating sciter ${SCITER_TAG} done; src at ${sciter_SOURCE_DIR}")
endif()

add_library(sciter INTERFACE)

# sciter end -------------------------------------------------------------------

# gRPC -------------------------------------------------------------------------

set(ABSL_PROPAGATE_CXX_STD ON)

set(GRPC_TAG "v1.49.1")
set(GRPC_SRC "${PROJECT_SOURCE_DIR}/../grpc-${GRPC_TAG}")

FetchContent_Declare(gRPC
        GIT_REPOSITORY  https://github.com/grpc/grpc.git
        GIT_TAG         ${GRPC_TAG}
        GIT_SHALLOW     TRUE
        SOURCE_DIR      ${GRPC_SRC}
        GIT_PROGRESS    TRUE
        )

FetchContent_MakeAvailable(gRPC)

# gRPC end ---------------------------------------------------------------------

# generate proto ---------------------------------------------------------------

#find_program(GRPC_CPP_PLUGIN grpc_cpp_plugin)
#find_program(PROTO_COMPILER protoc)
set(GRPC_CPP_PLUGIN "/home/catalin/.grpc/bin/grpc_cpp_plugin")
set(PROTO_COMPILER "/home/catalin/.grpc/bin/protoc")
message("GRPC_CPP_PLUGIN: ${GRPC_CPP_PLUGIN}")
message("PROTO_COMPILER: ${PROTO_COMPILER}")

get_filename_component(SENDER_PROTO_FILE ${PROJECT_SOURCE_DIR}/../proto/Sender/Sender.proto ABSOLUTE)
get_filename_component(SENDER_PROTO_PATH ${SENDER_PROTO_FILE} PATH)

#protoc --proto_path=../proto/Sender/ \
#    --cpp_out=. \
#    --grpc_out=. \
#    --plugin=protoc-gen-grpc=grpc_cpp_plugin \
#    ../proto/Sender/Sender.proto

execute_process(
        COMMAND
        ${PROTO_COMPILER}
        "--proto_path=${SENDER_PROTO_PATH}"
        "--cpp_out=${PROJECT_SOURCE_DIR}"
        "--grpc_out=${CMAKE_SOURCE_DIR}"
        "--plugin=protoc-gen-grpc=${GRPC_CPP_PLUGIN}"
        ${SENDER_PROTO_FILE}
)

# generate proto end -----------------------------------------------------------

# Setup CMake to use GTK+, tell the compiler where to look for headers
# and to the linker where to look for libraries
include_directories(${GTK3_INCLUDE_DIRS})

link_directories(${GTK3_LIBRARY_DIRS})

# Add other flags to the compiler
add_definitions(${GTK3_CFLAGS_OTHER})

include_directories("${SCITER_SDK}/include")

execute_process(
        COMMAND
            "${SCITER_SDK}/bin/linux/packfolder"
            ${PROJECT_SOURCE_DIR}/ui
            ${PROJECT_SOURCE_DIR}/resources.cpp
            "-v \"resources\""
        OUTPUT GENERATED_RESOURCES_FILE
)

set(HEADERS
        MainFrame.h
        ISender.h
        SocketSender.h
        RPCSender.h
        Sender.pb.h
        Sender.grpc.pb.h
        )

set(SOURCES
        # here is main, which calls uimain from main.cpp
        "${SCITER_SDK}/include/sciter-gtk-main.cpp"
        main.cpp
        MainFrame.cpp
        SocketSender.cpp
        RPCSender.cpp
        Sender.pb.cc
        Sender.grpc.pb.cc
        )

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
        ${LIBRARY}
        ${GENERATED_RESOURCES_FILE}
        ${COMPILED_PROTO_FILE}
        )

# Link the target to the GTK+ libraries
target_link_libraries(${PROJECT_NAME} ${GTK3_LIBRARIES} "${SCITER_SDK}/bin/linux/x64/libsciter.so")
target_link_libraries(${PROJECT_NAME} asio)
target_link_libraries(${PROJECT_NAME} grpc++)

configure_file("${SCITER_SDK}/bin/linux/x64/libsciter.so" ${CMAKE_BINARY_DIR}/libsciter.so COPYONLY)
