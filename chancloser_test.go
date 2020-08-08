package chancloser

import (
  "fmt"
  "runtime"
  "sync"
  "testing"
)

var running sync.WaitGroup

func sender(ch chan int){
  running.Add(1)
  defer running.Done()
  ChanClaim(ch)
  defer ChanRelease(ch)
  for i:=0;i<100;i++{
    ch <- i
  }
}

func receiver(ch chan int){
  running.Add(1)
  defer running.Done()
  for i:=range ch{
    fmt.Printf("got %d\n",i)
  }
}

func TestChanCloser(t *testing.T) {
  ch:=make(chan int,50)
  go sender(ch)
  go sender(ch)
  go receiver(ch)
  runtime.Gosched()
  running.Wait()
}
