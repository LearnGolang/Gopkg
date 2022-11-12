/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:2022/11/12 19:01
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func wgProcess(wg *sync.WaitGroup, id int) {
	fmt.Printf("process:%d is going!\n", id)
	wg.Done()
}

func main() {
	StartTime := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go wgProcess(wg, i)
	}
	wg.Wait()
	EndTime := time.Now()
	Time := EndTime.Sub(StartTime)
	fmt.Println(Time)
}
