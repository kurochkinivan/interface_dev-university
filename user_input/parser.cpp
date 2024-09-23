#include <string>
#include <algorithm>
#include <iostream>
#include "structs.h"

using namespace std;

string removeSpaces(const string &input)
{
    string result = input;
    result.erase(remove_if(result.begin(), result.end(), ::isspace), result.end());
    return result;
}

bool isPassportValid(const Passport &passport)
{
    if (passport.ss.length() != 4 || passport.nn.length() != 6)
        return false;
    
    for (int i = 0; i < passport.ss.length(); i++)
    {
        if (!isdigit(passport.ss[i]))
            return false;
    }

    for (int i = 0; i < passport.nn.length(); i++)
    {
        if (!isdigit(passport.nn[i]))
            return false;
    }

    return true;
}

bool isDateValid(string &date)
{
    date = removeSpaces(date);

    if (date.length() != 10)
        return false;

    for (int i = 0; i < date.length(); i++)
    {
        if (i == 2 || i == 5)
        {
            if (date[i] != '.')
                return false;
        }
        else
        {
            if (!isdigit(date[i]))
                return false;
        }
    }

    int dd = stoi(date.substr(0, 2));
    if (dd < 1 || 31 < dd)
        return false;

    int mm = stoi(date.substr(3, 2));
    if (mm < 1 || 12 < mm)
        return false;

    int yyyy = stoi(date.substr(6, 4));
    if (yyyy < 1900 || 2024 < yyyy)
        return false;

    return true;
}

bool isPhoneValid(string &phone)
{
    phone = removeSpaces(phone);

    if (!(phone.length() == 11 && phone[0] == '8') && !(phone.length() == 12 && phone[0] == '+' && phone[1] == '7'))
        return false;

    for (int i = 1; i < phone.length(); i++)
    {
        if (!isdigit(phone[i]))
            return false;
    }

    return true;
}

bool isTemperatureValid(double &temp)
{
    if (temp < 0 || temp > 50)
        return false;

    return true;
}

string formatPhone(string &phone)
{
    string formatted;
    if (phone[0] == '+')
    {
        formatted = phone.substr(0, 2) + "(" +
                    phone.substr(2, 3) + ") " +
                    phone.substr(5, 3) + "-" +
                    phone.substr(8, 2) + "-" +
                    phone.substr(10, 2);
        return formatted;
    }
    else if (phone[0] == '8')
    {
        formatted = phone.substr(0, 1) + "(" +
                    phone.substr(1, 3) + ") " +
                    phone.substr(4, 3) + "-" +
                    phone.substr(7, 4);
        return formatted;
    }

    return "";
}

Date parseDate(string &date)
{
    Date d;
    d.dd = stoi(date.substr(0, 2));
    d.mm = stoi(date.substr(3, 2));
    d.yyyy = stoi(date.substr(6, 4));
    return d;
}
