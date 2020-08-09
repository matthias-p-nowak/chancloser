package chancloser

import (
  "reflect"
  "sync"
  "log"
)

var allClosers map[interface{}]int 
var closerLock sync.Mutex

/*
Claims a channel for the Channel closer.
 ch is a channel of any kind
 Usage:
   ChanClaim(ch)
   defer ChanRelease(ch)
*/
func ChanClaim(ch interface{}){
  closerLock.Lock()
  defer closerLock.Unlock()  
  if allClosers[ch]<0 {
    log.Panic("Claiming channel after it was already closed")
  }
  allClosers[ch]++
  // fmt.Printf("allClosers is %#v\n",allClosers)
}

/*
Releases the claim on the channel, if all current claims have been released, it will close the channel. 
This function call should be a deferred call right after making the claim. 
 */
func ChanRelease(ch interface{}){
  closerLock.Lock()
  defer closerLock.Unlock()
  // fmt.Println("closing channel")
  allClosers[ch]-=1
  switch{
    case allClosers[ch]<0:
      log.Panic("another channel claimer closed channel erlier")
    case allClosers[ch]==0: 
      vo:=reflect.ValueOf(ch)
      vo=reflect.Indirect(vo)
      vo.Close()
      allClosers[ch]-=1
  }
}

func init(){
  allClosers=make(map[interface{}]int)
  // fmt.Println("chancloser initialized")
}
