//Smallest multiple
//
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

// Runtime 29.511s

#include <stdio.h>

int check(int n)
{
    int dent = 0;
    int sum = 1; int divisible = 20; 
    for (int i = 1;i <= divisible;i++) 
    {
        dent += n % i;
    }
    if (dent == 0) 
    {   
        return 1;
    }
    return 0;
}

int main () 
{
    int n = 1;
    while (1)
    {
        if (check(n) == 1) 
        {
            printf("%d",n);
            break;
        }
        n++;
    }
}
