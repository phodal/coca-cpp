# CMake 最低版本号要求
cmake_minimum_required(VERSION 3.15.0)

# 项目名
project(hello)

# 查找当前目录下的所有源文件，并将名称保存到 SRC_LIST 变量
aux_source_directory(. SRC_LIST)

# 指定生成目标
add_executable(hello ${SRC_LIST} line.h)
