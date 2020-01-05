package main

import (
	"fmt"
	"time"
)

func main() {
	timemap := make(map[string]*time.Timer)
	fmt.Print("\n1. adding a key ")
	timemap["name1"] = time.NewTimer(time.Second * 10)
	fmt.Print("\n 2. adding second key..")
	timemap["name2"] = time.NewTimer(time.Second * 10)
	fmt.Print("\n adding third key..")
	timemap["name3"] = time.NewTimer(time.Second * 10)

	fmt.Print("\n length of map is ", len(timemap))

	time.Sleep(time.Second * 5)
	fmt.Print("\n length of map is ", len(timemap))
	time.Sleep(time.Second * 5)
	fmt.Print("\n length of map is ", len(timemap))
}
