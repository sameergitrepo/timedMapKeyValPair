package main

import (
	"fmt"
	"time"
)

type TimeMap map[string]*time.Timer

func main() {
	timemap := make(TimeMap, 0)

	fmt.Print("\n1. adding a key ")
	timemap["name1"] = time.NewTimer(time.Second * 10)
	go timemap.Delete(timemap["name1"].C, "name1")
	fmt.Print("\n 2. adding second key..")
	timemap["name2"] = time.NewTimer(time.Second * 10)
	fmt.Print("\n adding third key..")
	timemap["name3"] = time.NewTimer(time.Second * 10)

	fmt.Print("\n length of map is ", len(timemap))

	time.Sleep(time.Second * 10)
	fmt.Print("\n length of map is ", len(timemap))
	time.Sleep(time.Second * 5)
	fmt.Print("\n length of map is ", len(timemap))
}

func (m TimeMap) Delete(timech <-chan time.Time, key string) {
	<-timech
	delete(m, key)
}

func (m TimeMap) AddToMap(key string, timervalinsec int) {
	m[key] = time.NewTimer(time.Second * time.Duration(timervalinsec))
	go m.Delete(m[key].C, key)
}
