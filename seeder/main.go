package main

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker"
)

func main() {
	nums, err := faker.RandomInt(1, 20000)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nums {
		fmt.Println(n)
	}
}
