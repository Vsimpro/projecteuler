// Sum square difference
// The sum of the squares of the first ten natural numbers is,
//                  1^2 + 2^2 + ... +10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
//                  (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers
// and the square of the sum is 3025 - 385 = 2640
// Find the difference between 
// the sum of the squares of the first one hundred natural numbers and the square of the sum.

// runtime  0.194 s

#include <stdio.h>

int main() 
{
    int difference = 0;
    int natural_numbers = 100;
    int sq_sum = 0; int sum_sq = 0;
    for (int i = 1; i <= natural_numbers;i++) 
    {
        sq_sum += i;
        sum_sq += i * i;
    } 
    sq_sum = sq_sum * sq_sum;
    printf("Square of sum: %d ",sq_sum);printf(" Sum of Square: %d\n",sum_sq);
    difference = sq_sum - sum_sq;
    printf("The difference between the sum of the squares of the first %d natural numbers,", natural_numbers);
    printf("\n and the square of the sum is %d",difference);
}
