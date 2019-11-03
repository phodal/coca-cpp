#include <iostream>
#include "line.h"


using namespace std;

// 成员函数定义，包括构造函数
Line::Line(void) {
    cout << "Object is being created" << endl;
}

void Line::setLength(double len) {
    length = len;
}

double Line::getLength(void) {
    return length;
}

