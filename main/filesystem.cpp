#include <iostream>
#include <fstream>
#include <cstring>
#include <unistd.h>
#include <sys/stat.h>

void Test01();
void Test02();
void Test03();
void Test04();
void Test05();

inline bool exists_test (const std::string& name) {
  struct stat buffer;   
  return (stat (name.c_str(), &buffer) == 0); 
}

int main() {
    Test05();
    return 0;
}

void Test01() {
    std::fstream iofile;  
    iofile.open("test1_cpp.txt", std::ios::out | std::ios::trunc);
    if (!iofile.is_open()){
        std::cout << "File Read Open" << std::endl;
        return;
    }
    char buf[4096];
    char data[100] = "Hello Cpp file io system!\n";
    for (int i = 0; i < 10; i++) {
        iofile<<data;
    }
    iofile.close();
    return;
}

void Test02() {
    std::fstream iofile;  
    iofile.open("test1_cpp.txt", std::ios::out | std::ios::app);
    if (!iofile.is_open()){
        std::cout << "File Read Open" << std::endl;
        return;
    }
    char data[100] = "ABC!ENGLISH!\n";
    iofile << data;
    iofile.close();
    return;
}

void Test03() {
    // 读取文件
    std::fstream iofile;  
    iofile.open("test1_cpp.txt", std::ios::in);
    if (!iofile.is_open()){
        std::cout << "File Read Open" << std::endl;
        return;
    }
   
    char buf[1024];
    while (!iofile.eof() ){
        iofile.getline(buf,1024);
        std::cout << buf << std::endl;  
    }
    iofile.close();

    // 后续追加输出
    iofile.open("test1_cpp.txt", std::ios::out | std::ios::app );
    std::string data= "This line is written by [cpp]\n";
    for (int i = 0; i < 5; i++ ){
        iofile << data << std::flush;
    }
    iofile.close();
    return;
}

void Test04() {
    // infile
    std::ifstream ifile;  
    ifile.open("test1_cpp.txt", std::ios::in);
    if (!ifile.is_open()){
        std::cout << "File Read Open" << std::endl;
        return;
    }
    // outfile
    std::ofstream ofile; 
    ofile.open("test1_cpp_1.txt", std::ios::out | std::ios::trunc );

   
    char buf[1024];
    while (!ifile.eof() ){
        ifile.getline(buf,1024);
        ofile << buf;
        ofile << "\n";
    }
    ifile.close();
    ofile.close();
    return;
}


void Test05() {
    std::ifstream ifile;
    ifile.open("test1_cpp.txt", std::ios::in);
    char buf[4096];
    int alphas = 0;
    int nums = 0;
    int spaces = 0;
    int others = 0;
    while (!ifile.eof()){
        ifile.getline(buf, 4096);
        int len = strlen(buf);
        for (int i = 0; i < len; i++ ){
            if (isalpha(buf[i])){
                alphas++;
            }else if (isdigit(buf[i])){
                nums ++;
            }else if (isspace(buf[i])){
                spaces ++;
            }else{
                others ++;
            }
        }
    }
    std::cout << "alphas: " << alphas << "\t" <<"nums: " 
                << nums << "\t" << "spaces: " << spaces << "\t" << "others: " << others << std::endl;
}



