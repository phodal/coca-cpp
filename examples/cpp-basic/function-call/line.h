//
// Created by Fengda Huang  on 2019/11/3.
//

#ifndef LINE_H
#define LINE_H

class Line {
public:
    void setLength(double len);

    double getLength(void);

    Line();  // 这是构造函数

private:
    double length;
};

#endif LINE_H
