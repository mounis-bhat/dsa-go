package main

import (
	"fmt"

	"github.com/mounis-bhat/dsa-go/internal/neetcode150"
)

func main() {
	res := neetcode150.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(res)
}
