package chancloser

import (
  "reflect"
  "sync"
)

var allClosers map[interface{}]int=make(map[interface{}]int)
var closerLock sync.Mutex

func ChanClaim(ch interface{}){
  closerLock.Lock()
  defer closerLock.Unlock()
  
  allClosers[ch]++
}

func ChanRelease(ch interface{}){
  closerLock.Lock()
  defer closerLock.Unlock()
  allClosers[ch]-=1
  if allClosers[ch] == 0 {
    vo:=reflect.ValueOf(ch)
    vo=reflect.Indirect(vo)
    vo.Close()
    delete(allClosers,ch)
  }
}
