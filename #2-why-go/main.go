/*This tiny Go program is doing real concurrency in just FOUR lines.

    A goroutine sends a message, a channel receives it,
    and boom, safe parallel execution without threads, locks, or headaches.

    This looks simple… but it’s actually powerful.
    . go func() creates a lightweight goroutine — like a super-cheap thread that Go manages for you.
    . ch := make(chan string) creates a channel… think of it like a safe pipe where goroutines send data without conflicts.
    . So one goroutine shouts ‘hello’ into the channel… and the main routine waits and receives it.
    . No race conditions.
    . No locks.
    . No complexity.
    . Real concurrency in 4 lines.
    Run this once with `go run main.go` — Once you see it, you’ll understand why backend + observability infra teams love Go.*/

  // watch https://youtube.com/shorts/iIag1z2AcSs?feature=share

    package main

func main() {
    ch := make(chan string)          
    go func() { ch <- "hello" }()
    msg := <-ch
    println(msg)
}

// Output: hello