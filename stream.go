package main

import (
	"log"
  "encoding/json"
  "fmt"

  "strconv"
	"github.com/gorilla/websocket"
)

type Streamer interface {
  Stream()
}




type OkxStream struct {
  conn *websocket.Conn
  ch chan Trade

}

func NewOkxStream(url string,ch chan Trade) *OkxStream {
  conn, _ ,err := websocket.DefaultDialer.Dial(url, nil)




  if err != nil {
    log.Fatalf("Error: %s", err)
  }
  return &OkxStream{
    conn: conn,
    ch: ch,
  }
}



func (s *OkxStream) Stream() {
  // TODO implement subscribtion to Okx
  defer s.conn.Close()

  fmt.Println("Stream started")
  instID := "BTC-USDT"
  channel := "trades"
  args := make([]OkxSubReqArg, 1)
  args[0] = OkxSubReqArg{channel, instID}

  okxSubReq := OkxSubReq{
    Op: "subscribe",
    OkxSubReqArgs: args,

  }

  

  subReq, err := json.Marshal(okxSubReq)
  if err != nil {
    log.Fatalf("Error: %s", err)
  }

  

  err = s.conn.WriteJSON(okxSubReq)
  if err != nil {
    log.Fatalf("Error: %s", err)
  }

  fmt.Printf("Subscribe request sent %s\n", subReq)

  subResp := new(OkxSubResp)

  // err = s.conn.ReadJSON(&subResp)
  _,b,err := s.conn.ReadMessage()
  if err != nil {

    log.Fatalf("Error: %s", err)
  }
  
  fmt.Println(err)
  err = json.Unmarshal(b, subResp)
  if err != nil {
    okxRespFail := OkxRespFail{}
    json.Unmarshal(b, &okxRespFail)
    log.Fatalf("Error: %s: %v", err,okxRespFail)
  }
  fmt.Println(err)

  fmt.Printf("Subscribe resp recv: %s\n", subResp)
  ltPush := LastTradePush{}
  for {
    err := s.conn.ReadJSON(&ltPush)
    if err != nil {
      log.Fatalf("Error: %s", err)
    }
    // trades := make([]Trade, len(ltPush.Data)) 
    for _, trade := range ltPush.Data {
      price, err := strconv.ParseFloat(trade.Px, 64)
      if err != nil {
        log.Fatalf("Error: %s", err)
      }
      volume, err := strconv.ParseFloat(trade.Sz, 64)
      if err != nil {
        log.Fatalf("Error: %s", err)
      }
      s.ch <- Trade{
        price: price,
        volume: volume,
      }
    }


    
  }

  





  
  
}
