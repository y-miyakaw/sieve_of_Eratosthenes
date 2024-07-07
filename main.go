package main

import (
	"fmt"
	"math"
)

func main() {
	primes := Sieve(1000)
	fmt.Println(primes)
}

func Sieve(max int) []int {
	if max < 2 {
		return []int{}
	}

	// 走査するスライスを作成
	// Create a slice to track prime candidates
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}

	// maxの平方根を算出
	// Calculate the square root of max
	sqrtMax := int(math.Sqrt(float64(max)))

	// エラトステネスの篩の走査を実行
	// Perform the Sieve of Eratosthenes
	for i := 2; i <= sqrtMax; i++ {
		if isPrime[i] {
			for j := i * i; j <= max; j += i {
				isPrime[j] = false
			}
		}
	}

	// エラトステネスの篩の走査によってtrueとなった数字を素数のスライスに格納
	// Store numbers that remain true after sieving in the primes slice
	primes := make([]int, 0, max/10)
	for i := 2; i <= max; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
