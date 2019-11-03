#include <iostream>

using namespace std;

void swap(int, int);

class Position {
private:
    int x;
private:
    int y;

public:
    void setX(int x) {
        this->x = x;
    }

    void setY(int y) {
        this->y = y;
    }

    void setPosition(int x, int y) {
        this->setX(x);
        this->setY(y);
    }
};

int main() {
    int a = 3, b = 4;
    cout << "a = " << a << ", b = "
         << b << endl;
    swap(a, b);
    cout << "a = " << a << ", b = "
         << b << endl;

    Position position;
    position.setX(0);
    return 0;
}

void swap(int x, int y) {
    int t = x;
    x = y;
    y = t;
}
