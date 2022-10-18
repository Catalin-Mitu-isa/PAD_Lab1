cmake_minimum_required(VERSION 3.18)
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
#set(GRPC_SRC "${PROJECT_SOURCE_DIR}/grpc-${GRPC_TAG}")

FetchContent_Declare(gRPC
        GIT_REPOSITORY  https://github.com/grpc/grpc.git
        GIT_TAG         ${GRPC_TAG}
        GIT_SHALLOW     TRUE
#        SOURCE_DIR      ${GRPC_SRC}
        GIT_PROGRESS    TRUE
        )

message(STATUS "Fetching grpc")
FetchContent_MakeAvailable(gRPC)
message(STATUS "grpc fetch done")

# gRPC end ---------------------------------------------------------------------

# generate proto ---------------------------------------------------------------

set(GRPC_CPP_PLUGIN $<TARGET_FILE:grpc_cpp_plugin>)
set(PROTO_COMPILER $<TARGET_FILE:protoc>)

get_filename_component(SENDER_PROTO_FILE ${PROJECT_SOURCE_DIR}/../proto/Sender.proto ABSOLUTE)
get_filename_component(SENDER_PROTO_PATH ${SENDER_PROTO_FILE} PATH)

#/home/catalin/.grpc/bin/protoc --proto_path=../proto/ --cpp_out=. --grpc_out=. --plugin=protoc-gen-grpc=/home/catalin/.grpc/bin/grpc_cpp_plugin ../proto/Sender.proto

set(_gRPC_PROTOBUF_PROTOC_EXECUTABLE $<TARGET_FILE:protoc>)
set(_gRPC_PROTO_GENS_DIR ${CMAKE_BINARY_DIR}/proto)
file(MAKE_DIRECTORY ${_gRPC_PROTO_GENS_DIR})
set(_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR ${PROTO_FILES_PATH})

function(target_add_protobuf target)
    if(NOT TARGET ${target})
        message(FATAL_ERROR "Target ${target} doesn't exist")
    endif()
    if(NOT ARGN)
        message(SEND_ERROR "Error: PROTOBUF_GENERATE_GRPC_CPP() called without any proto files")
        return()
    endif()

    foreach(FIL ${ARGN})
        get_filename_component(ABS_FIL ${FIL} ABSOLUTE)
        get_filename_component(FIL_PATH ${ABS_FIL} PATH)
        get_filename_component(FIL_WE ${FIL} NAME_WE)
        file(RELATIVE_PATH REL_FIL ${CMAKE_CURRENT_SOURCE_DIR} ${ABS_FIL})
        get_filename_component(REL_DIR ${REL_FIL} DIRECTORY)
        if(NOT REL_DIR)
            set(RELFIL_WE "${FIL_WE}")
        else()
            set(RELFIL_WE "${REL_DIR}/${FIL_WE}")
        endif()
        message("REL_FIL = ${REL_FIL}")
        message("ABS_FIL = ${ABS_FIL}")
        message("FIL_PATH = ${FIL_PATH}")

        if(NOT TARGET grpc_cpp_plugin)
            message(FATAL_ERROR "Can not find target grpc_cpp_plugin")
        endif()
        set(_gRPC_CPP_PLUGIN $<TARGET_FILE:grpc_cpp_plugin>)

        add_custom_command(
                OUTPUT  "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.cc"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.h"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}_mock.grpc.pb.h"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.cc"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.h"
                COMMAND ${_gRPC_PROTOBUF_PROTOC_EXECUTABLE}
                ARGS --grpc_out=${_gRPC_PROTO_GENS_DIR}
                --cpp_out=${_gRPC_PROTO_GENS_DIR}
                --plugin=protoc-gen-grpc=${_gRPC_CPP_PLUGIN}
                --proto_path=/home/catalin/UTM/Sem7/PAD/PAD_Lab1/Project/proto/
                ${ABS_FIL}
                DEPENDS ${ABS_FIL} ${_gRPC_PROTOBUF_PROTOC} ${_gRPC_CPP_PLUGIN}
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                COMMENT "Running gRPC C++ protocol buffer compiler on ${FIL}"
                VERBATIM)

        target_sources(${target} PRIVATE
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.cc"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.h"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.cc"
                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.h"
                )
        target_include_directories(${target} PRIVATE
                $<BUILD_INTERFACE:${_gRPC_PROTO_GENS_DIR}>
                $<BUILD_INTERFACE:${_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR}>
                $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/include>
                $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/third_party/abseil-cpp>
                )
    endforeach()
endfunction()

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
        )

set(SOURCES
        # here is main, which calls uimain from main.cpp
        "${SCITER_SDK}/include/sciter-gtk-main.cpp"
        main.cpp
        MainFrame.cpp
        SocketSender.cpp
        RPCSender.cpp
        )

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
        ${LIBRARY}
        ${GENERATED_RESOURCES_FILE}
        )

target_add_protobuf(${PROJECT_NAME} "../proto/Sender.proto")

# Link the target to the GTK+ libraries
target_link_libraries(${PROJECT_NAME} ${GTK3_LIBRARIES} "${SCITER_SDK}/bin/linux/x64/libsciter.so")
target_link_libraries(${PROJECT_NAME} asio)
target_link_libraries(${PROJECT_NAME} grpc++)

configure_file("${SCITER_SDK}/bin/linux/x64/libsciter.so" ${CMAKE_BINARY_DIR}/libsciter.so COPYONLY)
