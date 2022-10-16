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

# grpc + protobuf (copied) -----------------------------------------------------

set(CMAKE_POLICY_DEFAULT_CMP0077 NEW)
set(CMAKE_POLICY_DEFAULT_CMP0074 NEW)

set(FETCHCONTENT_QUIET OFF)
set(ABSL_PROPAGATE_CXX_STD ON)
set(ABSL_ENABLE_INSTALL ON)
#find_package(ZLIB REQUIRED)

#set(PROTOBUF_TAG "v3.19.3")
#set(PROTOBUF_SRC "${PROJECT_SOURCE_DIR}/../protobuf-${PROTOBUF_TAG}")

#FetchContent_Declare(
#        protobuf
#        GIT_REPOSITORY https://github.com/google/protobuf.git
#        GIT_TAG        ${PROTOBUF_TAG}
#        GIT_SHALLOW    TRUE
#        SOURCE_DIR     ${PROTOBUF_SRC}
#        GIT_PROGRESS   TRUE
#        USES_TERMINAL_DOWNLOAD TRUE
#        GIT_SUBMODULES_RECURSE FALSE
#        GIT_SUBMODULES ""
#)
#set(protobuf_BUILD_TESTS OFF)
#set(protobuf_BUILD_CONFORMANCE OFF)
#set(protobuf_BUILD_EXAMPLES OFF)
#set(protobuf_BUILD_PROTOC_BINARIES ON)
#set(protobuf_DISABLE_RTTI ON)
#set(protobuf_MSVC_STATIC_RUNTIME ON)
#set(protobuf_WITH_ZLIB ON CACHE BOOL "" FORCE)

#FetchContent_GetProperties(protobuf)
#if(NOT protobuf_POPULATED)
#    FetchContent_Populate(protobuf)
#    set(PROTOBUF_ROOT_DIR "${protobuf_SOURCE_DIR}")
#endif()

#FetchContent_GetProperties(protobuf)
#if(NOT protobuf_POPULATED)
#    message(STATUS "Fetching protobuf ${PROTOBUF_TAG}")
#
#    if(EXISTS "${PROTOBUF_SRC}")
#        set(FETCHCONTENT_SOURCE_DIR_PROTOBUF ${PROTOBUF_SRC})
#        message(STATUS "protobuf ${PROTOBUF_TAG} found; skip fetching.")
#    endif()
#
#    FetchContent_Populate(protobuf)
#    set(PROTOBUF_ROOT_DIR "${protobuf_SOURCE_DIR}")
#    message(STATUS "Populating protobuf ${PROTOBUF_TAG} done; src at ${protobuf_SOURCE_DIR}")
#endif()


# protobuf end

# grpc

set(GRPC_TAG "v1.49.1")
set(GRPC_SRC "${PROJECT_SOURCE_DIR}/../grpc-${GRPC_TAG}")

FetchContent_Declare(
        grpc
        GIT_REPOSITORY https://github.com/grpc/grpc.git
        GIT_TAG        ${GRPC_TAG}
        GIT_SHALLOW    TRUE
        SOURCE_DIR     ${GRPC_SRC}
        GIT_PROGRESS   TRUE
#        USES_TERMINAL_DOWNLOAD TRUE
#        GIT_SUBMODULES_RECURSE FALSE
#        GIT_SUBMODULES
#            "third_party/cares"
#            "third_party/boringssl-with-bazel"
#            "third_party/re2"
#            "third_party/abseil-cpp"
)

#set(gRPC_BUILD_TESTS OFF)
set(gRPC_BUILD_CODEGEN ON) # for grpc_cpp_plugin
#set(gRPC_BUILD_GRPC_CPP_PLUGIN ON) # we want to use only C++ plugin
#set(gRPC_BUILD_CSHARP_EXT OFF)
#set(gRPC_BUILD_GRPC_CSHARP_PLUGIN OFF)
#set(gRPC_BUILD_GRPC_NODE_PLUGIN OFF)
#set(gRPC_BUILD_GRPC_OBJECTIVE_C_PLUGIN OFF)
#set(gRPC_BUILD_GRPC_PHP_PLUGIN OFF)
#set(gRPC_BUILD_GRPC_PYTHON_PLUGIN OFF)
#set(gRPC_BUILD_GRPC_RUBY_PLUGIN OFF)

#set(gRPC_BENCHMARK_PROVIDER "none" CACHE STRING "" FORCE)
#set(gRPC_PROTOBUF_PROVIDER "module" CACHE STRING "" FORCE)
#set(gRPC_ZLIB_PROVIDER "package" CACHE STRING "" FORCE)

# use lite protobuf version, unless we start using features
# that require full protobuf
#set(gRPC_USE_PROTO_LITE ON CACHE BOOL "" FORCE)

#FetchContent_GetProperties(grpc)
#if(NOT grpc_POPULATED)
#    FetchContent_Populate(grpc)
#    add_subdirectory(${grpc_SOURCE_DIR} ${grpc_BINARY_DIR} EXCLUDE_FROM_ALL)
#endif()

FetchContent_GetProperties(grpc)
if(NOT grpc_POPULATED)
    message(STATUS "Fetching grpc ${GRPC_TAG}")

    if(EXISTS "${GRPC_SRC}")
        set(FETCHCONTENT_SOURCE_DIR_GRPC ${GRPC_SRC})
        message(STATUS "grpc ${GRPC_TAG} found; skip fetching.")
    endif()

    FetchContent_Populate(grpc)
    add_subdirectory(${grpc_SOURCE_DIR} ${grpc_BINARY_DIR} EXCLUDE_FROM_ALL)
    message(STATUS "Populating grpc ${GRPC_TAG} done; src at ${grpc_SOURCE_DIR}")
endif()

# grpc end

if(NOT TARGET protoc)
    message(FATAL_ERROR "Can not find target protoc")
endif()
set(_gRPC_PROTOBUF_PROTOC_EXECUTABLE $<TARGET_FILE:protoc>)
#message("???????_gRPC_PROTOBUF_PROTOC_EXECUTABLE = ${_gRPC_PROTOBUF_PROTOC_EXECUTABLE}")

if(NOT TARGET grpc_cpp_plugin)
    message(FATAL_ERROR "Can not find target grpc_cpp_plugin")
endif()
set(_gRPC_CPP_PLUGIN $<TARGET_FILE:grpc_cpp_plugin>)
#message("???????_gRPC_CPP_PLUGIN = ${_gRPC_CPP_PLUGIN}")

get_filename_component(SENDER_PROTO_FILE ${PROJECT_SOURCE_DIR}/../proto/Sender.proto ABSOLUTE)
get_filename_component(SENDER_PROTO_PATH ${SENDER_PROTO_FILE} PATH)

set(_gRPC_PROTO_GENS_DIR ${PROJECT_SOURCE_DIR})
#file(MAKE_DIRECTORY ${_gRPC_PROTO_GENS_DIR})

#execute_process(
#        COMMAND
#        ${_gRPC_PROTOBUF_PROTOC_EXECUTABLE}
#        "--proto_path=${SENDER_PROTO_PATH}"
#        "--cpp_out=${_gRPC_PROTO_GENS_DIR}"
#        "--grpc_out=${_gRPC_PROTO_GENS_DIR}"
#        "--plugin=protoc-gen-grpc=${_gRPC_CPP_PLUGIN}"
#        ${SENDER_PROTO_FILE}
#)
add_custom_command(
        OUTPUT
          "${_gRPC_PROTO_GENS_DIR}/Sender.grpc.pb.cc"
          "${_gRPC_PROTO_GENS_DIR}/Sender.grpc.pb.h"
          "${_gRPC_PROTO_GENS_DIR}/Sender.pb.cc"
          "${_gRPC_PROTO_GENS_DIR}/Sender.pb.h"
        COMMAND
          $<TARGET_FILE:protoc>
        ARGS
          --grpc_out=${_gRPC_PROTO_GENS_DIR}
          --cpp_out=${_gRPC_PROTO_GENS_DIR}
          --plugin=protoc-gen-grpc=$<TARGET_FILE:grpc_cpp_plugin>
          -I ${SENDER_PROTO_PATH}
        DEPENDS
          $<TARGET_NAME:protoc>
          $<TARGET_NAME:grpc_cpp_plugin>
#        WORKING_DIRECTORY
#          ${PROJECT_SOURCE_DIR}
)

set(GENERATED_PROTOS
        "${_gRPC_PROTO_GENS_DIR}/Sender.grpc.pb.cc"
        "${_gRPC_PROTO_GENS_DIR}/Sender.grpc.pb.h"
        "${_gRPC_PROTO_GENS_DIR}/Sender.pb.cc"
        "${_gRPC_PROTO_GENS_DIR}/Sender.pb.h"
)
#target_include_directories(${PROJECT_NAME}
#        ${_gRPC_PROTO_GENS_DIR}
#        $<BUILD_INTERFACE:${_gRPC_PROTO_GENS_DIR}>
#        $<BUILD_INTERFACE:${_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR}>
#        $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/include>
#        $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/third_party/abseil-cpp>
#        )

#set(_gRPC_PROTO_GENS_DIR ${CMAKE_BINARY_DIR}/gens)
#file(MAKE_DIRECTORY ${_gRPC_PROTO_GENS_DIR})
#set(_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR "${protobuf_SOURCE_DIR}/src")


# Example:
# In CMake:
#     add_library(tezt tezt.cpp)
#     target_add_protobuf(tezt my.proto)
# Then, in the source code:
#     #include <my.pb.h>
#     #include <my.grpc.pb.h>
#function(target_add_protobuf target)
#    if(NOT TARGET ${target})
#        message(FATAL_ERROR "Target ${target} doesn't exist")
#    endif()
#    if(NOT ARGN)
#        message(SEND_ERROR "Error: PROTOBUF_GENERATE_GRPC_CPP() called without any proto files")
#        return()
#    endif()

#    set(_protobuf_include_path -I . -I ${_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR})
#    foreach(FIL ${ARGN})
#        get_filename_component(ABS_FIL ${FIL} ABSOLUTE)
#        get_filename_component(FIL_WE ${FIL} NAME_WE)
#        file(RELATIVE_PATH REL_FIL ${CMAKE_CURRENT_SOURCE_DIR} ${ABS_FIL})
#        get_filename_component(REL_DIR ${REL_FIL} DIRECTORY)
#        if(NOT REL_DIR)
#            set(RELFIL_WE "${FIL_WE}")
#        else()
#            set(RELFIL_WE "${REL_DIR}/${FIL_WE}")
#        endif()

#        if(NOT TARGET grpc_cpp_plugin)
#            message(FATAL_ERROR "Can not find target grpc_cpp_plugin")
#        endif()
#        set(_gRPC_CPP_PLUGIN $<TARGET_FILE:grpc_cpp_plugin>)

#        add_custom_command(
#        OUTPUT  "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.cc"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.h"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}_mock.grpc.pb.h"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.cc"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.h"
#        COMMAND ${_gRPC_PROTOBUF_PROTOC_EXECUTABLE}
#        ARGS --grpc_out=generate_mock_code=true:${_gRPC_PROTO_GENS_DIR}
#            --cpp_out=${_gRPC_PROTO_GENS_DIR}
#            --plugin=protoc-gen-grpc=${_gRPC_CPP_PLUGIN}
#            ${_protobuf_include_path}
#            ${REL_FIL}
#        DEPENDS ${ABS_FIL} ${_gRPC_PROTOBUF_PROTOC} ${_gRPC_CPP_PLUGIN}
#        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
#        COMMENT "Running gRPC C++ protocol buffer compiler on ${FIL}"
#        VERBATIM)

#        execute_process(
#                COMMAND
#                ${PROTO_COMPILER}
#                "--proto_path=${SENDER_PROTO_PATH}"
#                "--cpp_out=${PROJECT_SOURCE_DIR}"
#                "--grpc_out=${CMAKE_SOURCE_DIR}"
#                "--plugin=protoc-gen-grpc=${GRPC_CPP_PLUGIN}"
#                ${SENDER_PROTO_FILE}
#        )

#        add_custom_command(
#        OUTPUT  "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.cc"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.h"
##                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}_mock.grpc.pb.h"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.cc"
#                "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.h"
#        COMMAND ${_gRPC_PROTOBUF_PROTOC_EXECUTABLE}
#        ARGS --grpc_out=${_gRPC_PROTO_GENS_DIR}
#            --cpp_out=${_gRPC_PROTO_GENS_DIR}
#            --plugin=protoc-gen-grpc=${_gRPC_CPP_PLUGIN}
#            ${_protobuf_include_path}
#            ${REL_FIL}
#        DEPENDS ${ABS_FIL} ${_gRPC_PROTOBUF_PROTOC} ${_gRPC_CPP_PLUGIN}
#        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
#        COMMENT "Running gRPC C++ protocol buffer compiler on ${FIL}"
#        VERBATIM
#        )


#                target_sources(${target} PRIVATE
#            "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.cc"
#            "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.grpc.pb.h"
##            "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}_mock.grpc.pb.h"
#            "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.cc"
#            "${_gRPC_PROTO_GENS_DIR}/${RELFIL_WE}.pb.h"
#        )
#        target_include_directories(${target} PRIVATE
#            $<BUILD_INTERFACE:${_gRPC_PROTO_GENS_DIR}>
#            $<BUILD_INTERFACE:${_gRPC_PROTOBUF_WELLKNOWN_INCLUDE_DIR}>
#            $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/include>
#            $<BUILD_INTERFACE:${grpc_SOURCE_DIR}/third_party/abseil-cpp>
#        )
#    endforeach()
#endfunction()

# grpc + protobuf end (copied) -------------------------------------------------

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
#        Sender.pb.h
#        Sender.grpc.pb.h
        )

set(SOURCES
        # here is main, which calls uimain from main.cpp
        "${SCITER_SDK}/include/sciter-gtk-main.cpp"
        main.cpp
        MainFrame.cpp
        SocketSender.cpp
        RPCSender.cpp
        # Sender.pb.cc
        # Sender.grpc.pb.cc
        )

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
        ${LIBRARY}
        ${GENERATED_RESOURCES_FILE}
        ${GENERATED_PROTOS}
        )

# Link the target to the GTK+ libraries
target_link_libraries(${PROJECT_NAME} ${GTK3_LIBRARIES} "${SCITER_SDK}/bin/linux/x64/libsciter.so")
target_link_libraries(${PROJECT_NAME} asio)
target_link_libraries(${PROJECT_NAME} RPCSender)

#target_add_protobuf(${PROJECT_NAME} "${PROJECT_SOURCE_DIR}/proto/Sender.proto")

#add_library(RPCSender RPCSender.cpp)
#get_filename_component(SENDER_PROTO_FILE ${PROJECT_SOURCE_DIR}/../proto/Sender.proto ABSOLUTE)
#get_filename_component(SENDER_PROTO_FILE_PATH ${SENDER_PROTO_FILE} PATH)
#message("-----------------------------${SENDER_PROTO_FILE_PATH}")
#target_add_protobuf(${PROJECT_NAME} proto/Sender.proto)
#target_link_libraries(${PROJECT_NAME} RPCSender)

#target_link_libraries(${PROJECT_NAME} grpc++)
# target_link_libraries(${PROJECT_NAME} gRPC)

configure_file("${SCITER_SDK}/bin/linux/x64/libsciter.so" ${CMAKE_BINARY_DIR}/libsciter.so COPYONLY)


#https://gist.githubusercontent.com/Warchant/2b75f6f903b4e71bbadaa163a8f82dfd/raw/02465cb848f146ad3930ef75e627efe398532112/grpc.cmake