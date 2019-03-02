package main

import (
	"fmt"
	"math"
)

const s = "never change"

func main() {
	fmt.Println(s)

	const d = 50000000

	const n = 3e20 / d

	fmt.Println(d, int64(n))

	fmt.Println(math.Sin(d))
}
