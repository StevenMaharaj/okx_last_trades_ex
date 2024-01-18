package main

import "log"

func main() {
  c := make(chan Trade)
  url := "wss://ws.okx.com:8443/ws/v5/public"

  okxStream := NewOkxStream(url,c)
  go okxStream.Stream()

  for {
    select {
    case trade := <-c:
      log.Printf("Trade: %v", trade)
    }
  }
}
