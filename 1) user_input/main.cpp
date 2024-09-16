#include <iostream>
#include <limits>
#include "structs.h"
#include "input.cpp"

using namespace std;

int main()
{
    Patient patient;

    // patient.name = getName();

    patient.passport = getPassport();

    // Date date = getDate();
    // patient.birth_date = date;

    // patient.phone = getPhone();

    // patient.temperature = getTemperature();

    cout << "Имя: " << patient.name << endl;
    cout << "Паспорт: " << patient.passport.ss << "-" << patient.passport.nn << endl;
    cout << "Дата рождения: " << patient.birth_date.dd << "." << patient.birth_date.mm << "." << patient.birth_date.yyyy << endl;
    cout << "Номер телефона: " << patient.phone << endl;
    cout << "Температура: " << patient.temperature << endl;
}