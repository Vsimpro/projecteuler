//Largest prime factor
//
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

package main

import (
        "fmt"
        "math"
)

func sieve(n int) []int {
        numbers := make([]int, n)

        for i := 0; i < len(numbers); i++ {
                numbers[i] = i
        }

        for j := 0; j < len(numbers); j++ {
                if j < 2 {
                        continue
                }

                if numbers[j] == 0 {
                        continue
                }
                for a := j; a < len(numbers) - 1; a++ {
                        if a * j >= len(numbers) {
                                break
                        }
                        numbers[a * j] = 0
                }
        }

        return numbers

}

func main() {
        n := 600851475143.00
        k := math.Sqrt(n)
        primes := sieve(int(k))

        last := 0
        for i := 0; i < len(primes); i++ {
                if primes[i] == 0 { continue }

                if int(n) % primes[i] == 0 {
                        last = primes[i]
                }
        }

        fmt.Printf("Last factor: %d\n", last)
}
