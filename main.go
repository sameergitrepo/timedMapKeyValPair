package main

import (
	"fmt"
	"time"
)

type TimeMap map[string]*time.Timer

func main() {
	timemap := make(TimeMap, 0)

	fmt.Print("\n1. adding a key ")

	timemap.AddToMap("name1", 10)
	timemap.AddToMap("name2", 15)
	timemap.AddToMap("name3", 20)

	fmt.Print("\n length of map is ", len(timemap))
	fmt.Print("\nsleeping...")
	time.Sleep(time.Second * 10)
	fmt.Print("\n length of map after sleeping for 10 seconds is ..")
	fmt.Print("\n length of map is ", len(timemap))
	fmt.Print("\n length of map after sleeping for 5 seconds is ..")
	time.Sleep(time.Second * 5)
	fmt.Print("\n final length of map is ", len(timemap))

}

func (m TimeMap) Delete(timech <-chan time.Time, key string) {
	fmt.Print("\n blocked for value ", key)
	<-timech
	delete(m, key)
	fmt.Print("\n deleted key.. ", key)
}

func (m TimeMap) AddToMap(key string, timervalinsec int) {
	m[key] = time.NewTimer(time.Second * time.Duration(timervalinsec))
	go m.Delete(m[key].C, key)
}
