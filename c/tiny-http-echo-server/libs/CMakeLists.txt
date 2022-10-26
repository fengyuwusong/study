# 添加源文件目录
aux_source_directory(. DIR_LIB_SRCS)
# 添加库
add_library(core ${DIR_LIB_SRCS})
