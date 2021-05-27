#include<iostream>

using namespace std;

int main(void){
    int a, b;
    int n;
    int cnt = 1;
    cin >> n;
    for (int i = 1 ; i <= 1000000 ;i++){
        cnt = cnt + 1;
    }
    while (n--)
    {
        cin >> a >> b;
        cout << a+b << endl;
    }
   
    return 0;
}
