#include <exception>
#include <iostream>
#include <string>
using namespace std;

class MyException : public exception {
public:
    const char* what(){
        return "My Exception";
    }
};

class ReadingFailException : public exception {
public:
    const char* what(){
        return "Reading Fail";
    }
};

double division (int a, int b) throw(MyException)
{
   if( b == 0 )
   {
      throw MyException();
   }
   return (a/b);
}

int test(int n1, int n2){
    return n1 + n2;
}

int (*ptest) (int , int);


template <typename T>
int sum(T a){
    return a;
}

template <typename T, typename... args>
int sum(T a, args... re){
    return a + sum(re...);
}

int readingFile(const string& fileName){
    if (fileName != "config.ini"){
        throw ReadingFailException();
    }
    return 0;
}

typedef struct {
    string name;
    int age;
    string color;
}Cat;

int main(){
    // ptest = &test;
    // cout << (*ptest)(10, 20) << endl;
    // cout << sum(10, 201, 239, 31, 321, 43, 54, 31) << endl;
    // int x = 50;
    // int y = 0;
    // double z = 0;
    // try {
    //     z = division(x, y);
    //     cout << z << endl;
    // }catch (MyException& e) {
    //     cerr << e.what() << endl;
    // }

    // try {
    //     readingFile("config.init");

    // }catch (ReadingFailException& e){
    //     cerr << e.what() << endl;
    // }
    Cat c1;
    c1.age = 1;
    c1.color = "white milk";
    c1.name = "fufu";
    cout << c1.name << "\t" << c1.age << "\t" << c1.color << endl;
    
    return 0;
}

