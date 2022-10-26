cmake_minimum_required(VERSION 3.10)

# 调试模式
set(CMAKE_BUILD_TYPE "Debug")
add_definitions("-Wall -g")

#设置项目名称
project(tiny-http-echo-server)

set(TARGET_FILE "http-echo-server")
set(CMAKE_BUILD_TYPE Release)
set(CMAKE_C_STANDARD 17)

include(CheckIncludeFiles)
check_include_files("pthread.h;stdatomic.h;sys/socket.h;netinet/in.h;unistd.h" EDEPS)
if (EPTHREAD EQUAL 1)
    message(FATAL_ERROR "Necessary header files are not found!")
endif()

# for headers in "/libs" and other external installed packages.
include_directories(. /opt/rh/devtoolset-11/root/lib/gcc/x86_64-redhat-linux/11/include)

# load source files and sub-directories.
# 添加源文件目录
aux_source_directory(./src DIR_SRCS)
# 添加子项目
add_subdirectory(libs/)

# load packages.
find_package(uriparser 0.9.6 CONFIG REQUIRED char)

# for executable.
# 设置二进制目标文件的依赖
add_executable(${TARGET_FILE} ${DIR_SRCS})
# 设置需要链接的库；
target_link_libraries(${TARGET_FILE} PUBLIC core m pthread uriparser::uriparser)
