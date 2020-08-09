package chancloser

import (
  "fmt"
  "runtime"
  "sync"
  "testing"
)

var running sync.WaitGroup

// one of 2 sender
func sender(ch chan int){
  running.Add(1)
  defer running.Done()
  Claim(ch)
  defer Release(ch)
  for i:=0;i<100;i++{
    ch <- i
  }
}

// the only receiver
func receiver(ch chan int){
  running.Add(1)
  defer running.Done()
  // runs until ch is closed
  for i:=range ch{
    fmt.Printf("got %d\n",i)
  }
}

// testing
func TestChanCloser(t *testing.T) {
  ch:=make(chan int,50)
  // 2 sender
  go sender(ch)
  go sender(ch)
  // the only receiver
  go receiver(ch)
  // give goroutines a chance
  runtime.Gosched()
  running.Wait()
}
