package main

import (
	"SDAlgorithm/golang/mipq"
	"fmt"
	"math/rand"
)

func main() {
	mq := mipq.Create()
	for i := 0; i < 10; i++ {
		n := rand.Intn(99) + 1
		fmt.Println(n)
		mq.Insert(n, n)
	}
	fmt.Println("......................")
	for mipq.Length(mq) != 0 {
		fmt.Println(mq.ExtractMin())
	}
	mq.Print()

}
