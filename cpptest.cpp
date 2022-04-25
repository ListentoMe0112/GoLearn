#include <iostream>
#include <thread>
#include <mutex>
#include <string>
#include <fstream>
#include <unordered_map>
#include <vector>
#include <algorithm>
#include <memory>
#include <type_traits>
#include <cstring>

const std::string FILENAME = "numbers.txt";

void writer(int i){
    std::ofstream of;
    std::string filename = FILENAME + std::to_string(i);
    of.open(filename, std::ios::binary | std::ios::trunc);
    for (int i = 0; i < 1000; i++ ){
        of << std::to_string(rand()) << "\n";
    }
    of.close();
}

void reader(int i){
    std::vector<int> nums;
    std::ifstream ifile;
    std::string filename = FILENAME + std::to_string(i);
    ifile.open(filename, std::ios::binary);
    char buf[4096];
    for (int i = 0; i < 1000; i++ ){
        ifile.getline(buf, 4096);
        nums.push_back(std::atoi(buf));
    }
    std::sort(nums.begin(), nums.end());
    ifile.close();

    std::ofstream of;
    of.open(filename, std::ios::binary | std::ios::trunc);
    for (int i = 0; i < 1000; i++ ){
        of << std::to_string(nums[i]) << "\n";
    }
}
 
int main(int argc, char *argv[])
{
    std::vector<std::thread> threads;
    for (int i = 0; i < 10; i++){
        threads.push_back(std::thread(writer, i));
    }
    for (int i = 0; i < 10; i++){
        threads[i].join();
    }

    std::vector<std::thread>().swap(threads);

    std::cin.get();
    for (int i = 0; i < 10; i++){
        threads.push_back(std::thread(reader, i));
    }
    for (int i = 0; i < 10; i++){
        threads[i].join();
    }
    return 0;
}




