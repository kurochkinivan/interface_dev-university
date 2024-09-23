#include <iostream>
#include <sstream>
#include <regex>
#include <limits>
#include "parser.cpp"
#include "structs.h"

using namespace std;

string invalidInputMessageF(const string &format)
{
    ostringstream oss;
    oss << "Неверный ввод! Формат ввода: " << format;
    return oss.str();
}

string getName()
{
    cout << "Введите имя пациента..." << endl;
    string name;
    while (true)
    {
        cin >> name;

        regex pattern("^[A-Za-z'-]+$");
        if (regex_match(name, pattern))
            return name;
        else
        {
            cerr << invalidInputMessageF("допустимые символы: 'A-Za-z', тире и апостроф.") << endl;
            continue;
        }
    }
}

Passport getPassport()
{
    Passport passport;
    while (true)
    {
        cout << "Введите серию паспорта пациента (4 цифры)..." << endl;
        cin >> passport.ss;
        cout << "Введите номер паспорта пациента (6 цифр)..." << endl;
        cin >> passport.nn;

        cout << "Паспорт: " << passport.ss << "-" << passport.nn << endl;

        if (isPassportValid(passport))
            return passport;
        else
        {
            cerr << invalidInputMessageF("серия: 4 цифры, номер: 6 цифр") << endl;
            continue;
        }
    }
}

Date getDate()
{
    cout << "Введите дату рождения пациента в формате DD.MM.YYYY..." << endl;
    string birth_date;
    while (true)
    {
        cin >> birth_date;

        if (isDateValid(birth_date))
        {
            return parseDate(birth_date);
        }
        else
        {
            cerr << invalidInputMessageF("DD.MM.YYYY") << endl;
            continue;
        }
    }
}

double getTemperature()
{
    cout << "Введите температуру пациента в формате числа с плавающей точкой..." << endl;
    double number;
    while (true)
    {
        if (cin >> number)
            if (isTemperatureValid(number))
                return number;
            else
            {
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                cerr << invalidInputMessageF("температура должна быть в диапазоне от 0 до 50") << endl;
                continue;
            }
        else
        {
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cerr << invalidInputMessageF("число с плавающей точкой") << endl;
        }
    }
}

string getPhone()
{
    cin.ignore(numeric_limits<streamsize>::max(), '\n');
    cout << "Введите номер телефона пациента..." << endl;
    string phone;
    while (true)
    {
        getline(cin, phone);

        if (isPhoneValid(phone))
        {
            phone = formatPhone(phone);
            return phone;
        }
        else
        {
            cerr << invalidInputMessageF("+7 XXX XXX XX XX or 8 XXX XXX XXXX") << endl;
            continue;
        }
    }
}