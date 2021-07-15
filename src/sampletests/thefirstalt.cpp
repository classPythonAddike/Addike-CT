#include <iostream>
#include <string>

using namespace std;

int main() {
    string num;
    int num2;

    cin >> num;
    num2 = stoi(num);

    cout << num2 * (num2 - 1) / 2 << endl;
}