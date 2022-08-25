//Special Pythagorean triplet
//
//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//a2 + b2 = c2
//For example, 32 + 42 = 9 + 16 = 25 = 52.
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product a*b*c.

package main

import (
        "math"
)

func main() {
        sumcap := 1000.00

        a := 0
        b := 0
        c := 0
        var c_temp float64 = 0

        for i := 1; i < int(sumcap); i++ {
                for j := i + 1; j < int(sumcap); j++  {
                        c_temp = float64(i*i + j*j)
                        c = int(math.Sqrt(c_temp))

                        sum := j + i + c
                        if sum == int(sumcap)  {
                                if (i < j) && (j < c) && (i*i + j*j == c*c) { 
                                        a = i
                                        b = j
                                        println("Triplet a, b, c; ", a, b, c)
                                        println("Product; ", a * b * c)
                                }
                        }
                }
        }
}
