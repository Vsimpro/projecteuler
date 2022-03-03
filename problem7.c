//10001st prime
//
//By listing the first six prime numbers: 
//2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?

#include <stdio.h>

int algo(int n) 
{
    for (int i = 2; i < n; i++)
    {
        if (n % i == 0 && n != i){
            return 0;
        }
    }
    // Prime!
    return 1;
}

int main()
{
    int key = 100000;
    int i = 2;int count = 0;
    while (1)
    {
        if (algo(i) == 1)
        {
            count++;
        }
        if (count >= 10001) 
        {
            printf("%d",i);
            break;
        }
        i++;
    }
  return 0;
}
