package main

import (
	"SDAlgorithm/golang/mapq"
	"fmt"
	"math/rand"
)

func main() {
	mq := mapq.Create()
	for i := 0; i < 10; i++ {
		n := rand.Intn(99) + 1
		fmt.Println(n)
		mq.Insert(n, n)
	}
	fmt.Println("......................")
	mq.Print()
	mq.Sort()
	mq.Print()

	/*
		for mapq.Length(mq) != 0 {
			fmt.Println(mq.ExtractMax())
		}
	*/

}
