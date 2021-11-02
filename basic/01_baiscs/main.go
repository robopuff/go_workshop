package main

import (
	"fmt"

	"github.com/mytheresa/go-workshop/internal"
)

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(*internal.Fizzbuzz(i))
	}
}
