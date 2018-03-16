package main

import (
	"SDAlgorithm/golang/brtree"
	"fmt"
	"strconv"
	"time"
)

func main() {
	BRT := brtree.Create()

	for i := 0; i < 100000; i++ { //use b.N for looping
		brtree.Insert(BRT, strconv.Itoa(i), i)
	}
	A := time.Now().UnixNano()
	brtree.Delete(BRT, "51452")
	C := time.Now().UnixNano()
	fmt.Println(A, C)
	fmt.Println(C - A)

}
