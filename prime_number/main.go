// This program takes a number N as input and generates primes
// numbers upto that number.
package main

import (
    "fmt"
    "math/big"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)
    for i := 0; i < N; i++ {
        if big.NewInt(int64(i)).ProbablyPrime(1) {
            fmt.Println(i)
        }
    }
}
