cmake_minimum_required( VERSION 2.8)
set( CMAKE_CXX_FLAGS "-std=c++11" )
PROJECT(cpptest)

set(SRC 
    ${PROJECT_SOURCE_DIR}/cpptest.cpp 
)

include_directories(
${PROJECT_SOURCE_DIR}
)

add_executable(${PROJECT_NAME} ${SRC})

target_link_libraries(${PROJECT_NAME} pthread dl)