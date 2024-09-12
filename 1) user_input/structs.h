#ifndef STRUCTS_H
#define STRUCTS_H

#include <string>

using namespace std;

struct Date {
    int dd;
    int mm;
    int yyyy;
};

struct Patient {
    string passport;    // ss ss-nnnnnn
    string name;        // any string
    Date birth_date;    // yyyy-mm-dd
    string phone;       // +X(XXX) XXX-XX-XX or X(XXX) XXX-XXXX
    double temperature; // XX.XX
};

#endif
