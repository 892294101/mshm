package main

import (
	"fmt"
	"github.com/892294101/mshm"
	shmdata "github.com/892294101/mshm/ishm"
	"log"
)

func testConstructor() interface{} {
	return &shmdata.TagTLV{}
}

//todo run this please run del-shm.sh
func main() {

	mem, err := shm.NewSystemVMem(6, 10000)

	if err != nil {
		log.Fatal(err)
	}

	s, err := shm.NewMultiShm(mem, 10000, testConstructor)
	if err != nil {
		fmt.Println(err)
		return
	}

	items, err := s.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range items {
		fmt.Printf("value : %v, type = %T\n", v, v)
	}
}
