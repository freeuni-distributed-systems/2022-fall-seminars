//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, ch chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(ch)
			return
		}

		ch <- tweet
	}
}

func consumer(ch <-chan *Tweet, wg *sync.WaitGroup) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	wg.Done() // 0
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	ch := make(chan *Tweet, 10)
   var wg sync.WaitGroup // 0

   wg.Add(1) // 1
	// Producer
	go producer(stream, ch)

	// Consumer
	go consumer(ch, &wg)

	wg.Wait() //0


	fmt.Printf("Process took %s\n", time.Since(start))
}
