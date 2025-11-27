/*  
This tiny Go snippet shows you the REAL power of Go’s concurrency model.

Every time this loop runs, Go launches a new goroutine — a super-lightweight, 
almost-free thread managed by Go’s own scheduler.

Each goroutine prints its own number, and Go automatically spreads them across 
multiple CPU cores. No threads, no locks, no configuration, no headaches.  
Just `go func()` and it runs in parallel.

`fmt.Scanln()` at the end simply keeps the program alive long enough for all the 
goroutines to complete, otherwise they might exit too fast.

This example looks simple… but this is EXACTLY how huge cloud systems, 
observability pipelines, and distributed backends scale with Go.  
You start small → then scale to thousands of goroutines effortlessly.

🔥 Your Challenge:
Can you modify this program so each goroutine sleeps for a random duration 
and prints when it starts and ends?  
(You’ll instantly see how Go handles massive parallel work.)

watch https://www.youtube.com/shorts/2L_PM2ETioI
*/


package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        go func(n int) {
            fmt.Println("running:", n)
        }(i)
    }
    fmt.Scanln() // wait
}
