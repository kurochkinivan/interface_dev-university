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

        // TODO: Решить проблему с символами кириллицы
        regex pattern("^[A-Za-zА-Яа-яЁё' -]+$");
        if (regex_match(name, pattern))
            return name;
        else
        {
            cerr << invalidInputMessageF("допустимые символы: 'A-Za-z', тире и апостроф.") << endl;
            continue;
        }
    }
}

string getPassport()
{
    cout << "Введите паспорт пациента в формате (ssss-nnnnnn)..." << endl;
    string passport;
    while (true)
    {
        cin >> passport;

        if (isPassportValid(passport))
            return passport;
        else
        {
            cerr << invalidInputMessageF("ssss-nnnnnn (тире обязательно)") << endl;
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
            auto [dd, mm, yyyy] = parseDate(birth_date);
            Date d;
            d.dd = dd;
            d.mm = mm;
            d.yyyy = yyyy;
            return d;
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
            return number;
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