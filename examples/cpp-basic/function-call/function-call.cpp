#include <iostream>

using namespace std;

void swap(int, int);

int main() {
    int a = 3, b = 4;
    cout << "a = " << a << ", b = "
         << b << endl;
    swap(a, b);
    cout << "a = " << a << ", b = "
         << b << endl;
    return 0;
}

void swap(int x, int y) {
    int t = x;
    x = y;
    y = t;
}
