cmake_minimum_required(VERSION 3.22)
project(Sender)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Use the package PkgConfig to detect GTK+ headers/library files
find_package(PkgConfig REQUIRED)
pkg_check_modules(GTK3 REQUIRED gtk+-x11-3.0)

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
        Sender.h
        )

set(SOURCES
        # here is main, which calls uimain from main.cpp
        "${SCITER_SDK}/include/sciter-gtk-main.cpp"
        main.cpp
        MainFrame.cpp
        Sender.cpp
        ${GENERATED_RESOURCES_FILE}
        )

add_executable(Sender ${SOURCES} ${HEADERS} ${LIBRARY})

# Link the target to the GTK+ libraries
target_link_libraries(Sender ${GTK3_LIBRARIES} "${SCITER_SDK}/bin/linux/x64/libsciter.so")

configure_file("${SCITER_SDK}/bin/linux/x64/libsciter.so" ${CMAKE_BINARY_DIR}/libsciter.so COPYONLY)
