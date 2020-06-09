package main

import (
	"fmt"
	"sync"
	"time"
)

var reqque chan interface{}

const (
	rquesize = 1000 //Request que size
	tps      = 100  //TPS required
)

func init() {

	//Declaring the request que to 1000 initially
	reqque = make(chan interface{}, rquesize)
}

func main() {
	var wg sync.WaitGroup
	var cnt int

	//Starting the request generator
	wg.Add(1)
	go reqgenrator(&wg)
	for {

		if cnt == tps {
			fmt.Printf("\n%d Requests processed %s", cnt, time.Now())
			//fmt.Printf("%d Number of Active go routines", os.pro)

			time.Sleep(1 * time.Second) //Injecting  the delay to ensure throttling
			cnt = 0
		}

		wg.Add(1)
		go reqprocessor(&wg, <-reqque)

		cnt++

	}

	wg.Wait()

}

func reqprocessor(wg *sync.WaitGroup, req interface{}) {

	//fmt.Println("\nProcessing : ", req.(string)) //Type asserting as we are aware of the underlying type

	wg.Done()
}

func reqgenrator(wg *sync.WaitGroup) {
	var i int

	for {

		reqque <- fmt.Sprintf("Request  : %d", i)
		i++

		if i%500 == 0 {
			fmt.Printf("\n%d  Request Generated ..", i)
		}
	}
	wg.Done()

}
