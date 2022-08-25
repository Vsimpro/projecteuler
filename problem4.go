//Largest palindrome product
//
//A palindromic number reads the same both ways.
//The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
        "fmt"
        "strconv"
)

func reverseString(str string) (result string) {
        for _ , v := range str {
                result = string(v) + result
        }
        return
}

func main() {
        n :=  1000
        chars := ""

        lastPal := 0
        for i := 1; i < n; i++ {
                for j := 1; j < n; j++ {
                        chars = strconv.Itoa(j * i)
                        reverse := reverseString(chars)
                        if chars == reverse {
                                if lastPal < j * i {
                                        lastPal = j * i
                                }
                        }

                }

        }
        fmt.Printf("Largest palindrome: %d\n", lastPal)
}
