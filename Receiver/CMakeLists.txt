cmake_minimum_required(VERSION 3.22)
project(Receiver)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Use the package PkgConfig to detect GTK+ headers/library files
find_package(PkgConfig REQUIRED)
pkg_check_modules(GTK3 REQUIRED gtk+-x11-3.0)

# Asio -------------------------------------------------------------------------

set(ASIO_CMAKE_ASIO_TAG "asio-1-24-0")
set(ASIO_CMAKE_ASIO_SOURCE_DIR ${CMAKE_CURRENT_LIST_DIR})
set(ASIO_CMAKE_ASIO_DEP_DIR "${ASIO_CMAKE_ASIO_SOURCE_DIR}/${ASIO_CMAKE_ASIO_TAG}-src")

include(FetchContent)

FetchContent_Declare(asio
        GIT_REPOSITORY  https://github.com/chriskohlhoff/asio.git
        GIT_TAG         ${ASIO_CMAKE_ASIO_TAG}
        GIT_SHALLOW     TRUE
        SOURCE_DIR      ${ASIO_CMAKE_ASIO_DEP_DIR}
        )

FetchContent_GetProperties(asio)

if(NOT asio_POPULATED)
    message(STATUS "Fetching asio ${ASIO_CMAKE_ASIO_TAG}")

    if(EXISTS "${ASIO_CMAKE_ASIO_DEP_DIR}/asio")
        set(FETCHCONTENT_SOURCE_DIR_ASIO ${ASIO_CMAKE_ASIO_DEP_DIR})
        message(STATUS "asio ${ASIO_CMAKE_ASIO_TAG} source found at ${ASIO_CMAKE_ASIO_DEP_DIR}; skip fetching.")
    endif()

    FetchContent_Populate(asio)
    message(STATUS "Populating asio ${ASIO_CMAKE_ASIO_TAG} done; src at ${asio_SOURCE_DIR}")
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

FetchContent_Declare(sciter
        GIT_REPOSITORY  https://gitlab.com/sciter-engine/sciter-js-sdk.git
        GIT_TAG         "5.0.0.9"
        GIT_SHALLOW     TRUE
        SOURCE_DIR      "${CMAKE_CURRENT_LIST_DIR}/sciter-js-sdk"
        )

FetchContent_GetProperties(sciter)

if(NOT sciter_POPULATED)
    message(STATUS "Fetching sciter 5.0.0.9")

    if(EXISTS "${CMAKE_CURRENT_LIST_DIR}/sciter-js-sdk")
        set(FETCHCONTENT_SOURCE_DIR_SCITER "${CMAKE_CURRENT_LIST_DIR}/sciter-js-sdk")
        message(STATUS "sciter found; skip fetching.")
    endif()

    FetchContent_Populate(sciter)
    message(STATUS "Populating sciter done; src at ${sciter_SOURCE_DIR}")
endif()

add_library(sciter INTERFACE)

# sciter end -------------------------------------------------------------------

# Setup CMake to use GTK+, tell the compiler where to look for headers
# and to the linker where to look for libraries
include_directories(${GTK3_INCLUDE_DIRS})

link_directories(${GTK3_LIBRARY_DIRS})

# Add other flags to the compiler
add_definitions(${GTK3_CFLAGS_OTHER})

set(SCITER_SDK "${PROJECT_SOURCE_DIR}/sciter-js-sdk")

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
        IReceiver.h
        SocketReceiver.h
        RPCReceiver.h
        )

set(SOURCES
        # here is main, which calls uimain from main.cpp
        "${SCITER_SDK}/include/sciter-gtk-main.cpp"
        main.cpp
        MainFrame.cpp
        SocketReceiver.cpp
        RPCReceiver.cpp
        ${GENERATED_RESOURCES_FILE}
        )

add_executable(${PROJECT_NAME} ${SOURCES} ${HEADERS} ${LIBRARY})

# Link the target to the GTK+ libraries
target_link_libraries(${PROJECT_NAME} ${GTK3_LIBRARIES} "${SCITER_SDK}/bin/linux/x64/libsciter.so")
target_link_libraries(${PROJECT_NAME} asio)

configure_file("${SCITER_SDK}/bin/linux/x64/libsciter.so" ${CMAKE_BINARY_DIR}/libsciter.so COPYONLY)
